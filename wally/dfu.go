package wally

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/gousb"
)

type status struct {
	bStatus       string
	bwPollTimeout int
	bState        string
	iString       string
}

func dfuCommand(dev *gousb.Device, addr int, command int, status *status) (err error) {
	var buf []byte
	if command == setAddress {
		buf = make([]byte, 5)
		buf[0] = 0x21
		buf[1] = byte(addr & 0xff)
		buf[2] = byte((addr >> 8) & 0xff)
		buf[3] = byte((addr >> 16) & 0xff)
		buf[4] = byte((addr >> 24) & 0xff)
	}
	if command == eraseAddress {
		buf = make([]byte, 5)
		buf[0] = 0x41
		buf[1] = byte(addr & 0xff)
		buf[2] = byte((addr >> 8) & 0xff)
		buf[3] = byte((addr >> 16) & 0xff)
		buf[4] = byte((addr >> 24) & 0xff)
	}
	if command == eraseFlash {
		buf = make([]byte, 1)
		buf[0] = 0x41
	}

	_, err = dev.Control(33, 1, 0, 0, buf)

	err = dfuPollTimeout(dev, status)

	if err != nil {
		return err
	}

	return nil
}

func dfuPollTimeout(dev *gousb.Device, status *status) (err error) {
	for i := 0; i < 3; i++ {
		err = dfuGetStatus(dev, status)
		time.Sleep(time.Duration(status.bwPollTimeout) * time.Millisecond)
	}
	return err
}

func dfuGetStatus(dev *gousb.Device, status *status) (err error) {
	buf := make([]byte, 6)
	stat, err := dev.Control(161, 3, 0, 0, buf)
	if err != nil {
		return err
	}
	if stat == 6 {
		status.bStatus = string(buf[0])
		status.bwPollTimeout = int((0xff & buf[3] << 16) | (0xff & buf[2]) | 0xff&buf[1])
		status.bState = string(buf[4])
		status.iString = string(buf[5])
	}
	return err
}

func dfuClearStatus(dev *gousb.Device) (err error) {
	_, err = dev.Control(33, 4, 2, 0, nil)
	return err
}

func dfuReboot(dev *gousb.Device, status *status) (err error) {
	err = dfuPollTimeout(dev, status)
	_, err = dev.Control(33, 1, 2, 0, nil)
	time.Sleep(1000 * time.Millisecond)
	err = dfuGetStatus(dev, status)
	return err
}

func extractSuffix(fileData []byte) (hasSuffix bool, data []byte, err error) {

	fileSize := len(fileData)

	suffix := fileData[fileSize-dfuSuffixLength : fileSize]
	d := string(suffix[10])
	f := string(suffix[9])
	u := string(suffix[8])

	if d == "D" && f == "F" && u == "U" {
		vid := int((suffix[5] << 8) + suffix[4])
		pid := int((suffix[3] << 8) + suffix[2])
		if vid != dfuSuffixVendorID || pid != dfuSuffixProductID {
			message := fmt.Sprintf("Invalid vendor or product id, expected %#x:%#x got %#x:%#x", dfuSuffixVendorID, dfuSuffixProductID, vid, pid)
			err = errors.New(message)
			return true, fileData, err

		}

		return true, fileData[0 : fileSize-dfuSuffixLength], nil
	}

	return false, fileData, nil
}

func DFUFlash(s *State) {
	dfuStatus := status{}
	fileData, err := os.ReadFile(s.FirmwarePath)
	if err != nil {
		message := fmt.Sprintf("Error while opening firmware: %s", err)
		s.Log("error", message)
		return
	}

	hasSuffix, firmwareData, err := extractSuffix(fileData)
	if err != nil {
		message := fmt.Sprintf("Error while extracting DFU Suffix: %s", err)
		s.Log("error", message)
		return
	}
	if hasSuffix {
		s.Log("info", "Found a valid DFU Suffix")
	} else {
		s.Log("info", "No DFU Suffix found")
	}

	ctx := gousb.NewContext()
	defer ctx.Close()
	var dev *gousb.Device

	// Get the list of device that match TMK's vendor id
	for {
		// if the app is reset stop this goroutine and close the usb context
		if s.Step != 3 {
			s.Log("info", "App reset, interrupting the flashing process.")
			return
		}
		s.Log("info", "Waiting for a DFU capable device")
		devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
			if desc.Vendor == gousb.ID(dfuVendorID) && desc.Product == gousb.ID(dfuProductID) {
				return true
			}
			return false
		})

		defer func() {
			for _, d := range devs {
				d.Close()
			}
		}()

		if err != nil {
			message := fmt.Sprintf("OpenDevices: %s", err)
			s.Log("warning", message)
		}

		if len(devs) > 0 {
			dev = devs[0]
			break
		}
		time.Sleep(1 * time.Second)
	}

	dev.SetAutoDetach(true)

	dev.ControlTimeout = 5 * time.Second

	cfg, err := dev.Config(1)
	if err != nil {
		message := fmt.Sprintf("Error while claiming the usb interface: %s", err)
		s.Log("error", message)
		return
	}
	defer cfg.Close()

	fileSize := len(firmwareData)
	s.FlashProgress.Total = fileSize

	err = dfuClearStatus(dev)
	if err != nil {
		message := fmt.Sprintf("Error while clearing the device status: %s", err)
		s.Log("error", message)
		return
	}

	s.Step = 4
	s.emitUpdate()

	err = dfuCommand(dev, 0, eraseFlash, &dfuStatus)
	if err != nil {
		err = fmt.Errorf("Error while erasing flash:", err)
		return
	}

	for page := 0; page < fileSize; page += planckBlockSize {
		addr := planckStartAddress + page
		chunckSize := planckBlockSize

		if page+chunckSize > fileSize {
			chunckSize = fileSize - page
		}

		err = dfuCommand(dev, addr, eraseAddress, &dfuStatus)
		if err != nil {
			message := fmt.Sprintf("Error while sending the erase address command: %s", err)
			s.Log("error", message)
			return
		}
		err = dfuCommand(dev, addr, setAddress, &dfuStatus)
		if err != nil {
			message := fmt.Sprintf("Error while sending the set address command: %s", err)
			s.Log("error", message)
			return
		}

		buf := firmwareData[page : page+chunckSize]
		bytes, err := dev.Control(33, 1, 2, 0, buf)

		if err != nil {
			message := fmt.Sprintf("Error while sending firmware bytes: %s", err)
			s.Log("error", message)
			return
		}

		message := fmt.Sprintf("Sent %d bytes out of %d", page+bytes, fileSize)
		s.Log("info", message)
		s.FlashProgress.Sent += bytes
		s.emitUpdate()
	}

	time.Sleep(1 * time.Second)

	s.Log("info", "Sending the reboot command")

	err = dfuReboot(dev, &dfuStatus)
	if err != nil {
		message := fmt.Sprintf("Error while rebooting device: %s", err)
		s.Log("error", message)
		return
	}

	s.Step = 5
	s.Log("info", "Flash complete")
}
