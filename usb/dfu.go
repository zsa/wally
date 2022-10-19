package usb

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	//DFU DATA SETTINGS
	DFU_SUFFIX_LENGTH   = 16
	DFU_PACKET_RETRIES  = 5
	STM32_START_ADDRESS = 0x08000000
	STM32_END_ADDRESS   = 0x0803ffff

	//DFU Statuses
	DFU_OK       = 0x0A
	DFU_IDLE     = 0x02
	DFU_DNBUSY   = 0x04
	DFU_DNIDLE   = 0x05
	DFU_MANIFEST = 0x07
	DFU_ERROR    = 0x0A

	// DFU Class requests
	DFU_DETACH    = 0x00
	DFU_DNLOAD    = 0x01
	DFU_UPLOAD    = 0x02
	DFU_GETSTATUS = 0x03
	DFU_CLRSTATUS = 0x04
	DFU_GETSTATE  = 0x05
	DFU_ABORT     = 0x06

	//DFU Commands
	DFU_SET_ADDRESS  = 0x21
	DFU_ERASE_SECTOR = 0x41
)

type dfuStatus struct {
	bStatus       uint8
	bwPollTimeout uint32
	bState        uint8
}

/*DFU commands*/

func (d *USBDevice) dfuPollTimeout(predicate func() bool) (err error) {
	for predicate() {
		d.cb(FlashCallback{Message: fmt.Sprintf("[DFU] Waiting for device, State: 0x%X, Status: 0x%X, Timeout: %dms", d.flashStatus.bState, d.flashStatus.bStatus, d.flashStatus.bwPollTimeout), Type: Log})

		time.Sleep(time.Duration(d.flashStatus.bwPollTimeout) * time.Millisecond)

		err := d.dfuGetStatus()
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *USBDevice) dfuClearStatus() (err error) {
	_, err = d.Control(33, DFU_CLRSTATUS, 0, 0, nil)
	if err != nil {
		return err
	}
	err = d.dfuGetStatus()
	if err != nil {
		return err
	}
	if d.flashStatus.bStatus == DFU_ERROR {
		return errors.New("[DFU] Failed clearing status, try to unplug / plug the device and retry flashing")
	}
	return nil
}

func (d *USBDevice) dfuCommand(cmd byte, addr int) (err error) {

	var buf []byte
	if cmd == DFU_SET_ADDRESS {
		buf = make([]byte, 5)
		buf[0] = DFU_SET_ADDRESS
		buf[1] = byte(addr & 0xff)
		buf[2] = byte((addr >> 8) & 0xff)
		buf[3] = byte((addr >> 16) & 0xff)
		buf[4] = byte((addr >> 24) & 0xff)
	}
	if cmd == DFU_ERASE_SECTOR {
		if addr == 0 {
			buf = make([]byte, 1)
			buf[0] = DFU_ERASE_SECTOR
		} else {
			buf = make([]byte, 5)
			buf[0] = DFU_ERASE_SECTOR
			buf[1] = byte(addr & 0xff)
			buf[2] = byte((addr >> 8) & 0xff)
			buf[3] = byte((addr >> 16) & 0xff)
			buf[4] = byte((addr >> 24) & 0xff)
		}
	}
	_, err = d.Control(33, DFU_DNLOAD, 0, 0, buf)
	if err != nil {
		return err
	}

	err = d.dfuGetStatus()
	if err != nil {
		return err
	}

	err = d.dfuPollTimeout(func() bool { return d.flashStatus.bState == DFU_DNBUSY })
	if err != nil {
		return err
	}

	return nil
}

func (d *USBDevice) dfuDownload(data []byte) (err error) {
	_, err = d.Control(33, DFU_DNLOAD, 2, 0, data)
	if err != nil {
		return err
	}

	err = d.dfuGetStatus()

	if err != nil {
		return err
	}

	err = d.dfuPollTimeout(func() bool { return d.flashStatus.bState != DFU_DNIDLE })
	if err != nil {
		return err
	}

	return err
}

func (d *USBDevice) dfuGetStatus() (err error) {
	buf := make([]byte, 6)
	res, err := d.Control(161, DFU_GETSTATUS, 0, 0, buf)
	if err != nil {
		return err
	}
	d.flashStatus.bStatus = res[0]
	d.flashStatus.bwPollTimeout = uint32(res[1]) | uint32(res[2])<<8 | uint32(res[3])<<16
	d.flashStatus.bState = res[4]
	return nil
}

func (d *USBDevice) dfuReboot() (err error) {
	_, err = d.Control(33, DFU_DNLOAD, 2, 0, nil)
	if err != nil {
		return err
	}

	err = d.dfuGetStatus()
	if err != nil {
		return err
	}

	err = d.dfuPollTimeout(func() bool { return d.flashStatus.bState != DFU_MANIFEST })

	return err
}

func (d *USBDevice) dfuEraseSectors(start, end, blockSize int) (err error) {

	for addr := start; addr < end; addr += blockSize {
		fmt.Printf("Erasing sector 0x%X", addr)
		err = d.dfuCommand(DFU_ERASE_SECTOR, addr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *USBDevice) dfuCheckState() (err error) {
	if d.flashStatus.bStatus != DFU_OK {
		return fmt.Errorf("device state error, bStatus 0x%X, bState 0x%X", d.flashStatus.bStatus, d.flashStatus.bState)
	}

	return nil
}

/* DFU Utils */
func (dev *USBDevice) extractSuffix(fileData []byte) (hasSuffix bool, data []byte, err error) {

	fileSize := len(fileData)
	vid := dev.Handle.GetVid()
	pid := dev.Handle.GetPid()

	suffix := fileData[fileSize-DFU_SUFFIX_LENGTH : fileSize]
	d := string(suffix[10])
	f := string(suffix[9])
	u := string(suffix[8])

	if d == "D" && f == "F" && u == "U" {
		fvid := int(suffix[5])<<8 | int(suffix[4])
		fpid := int(suffix[3])<<8 | int(suffix[2])
		if vid != fvid || pid != fpid {
			message := fmt.Sprintf("Invalid vendor or product id, expected %#x:%#x got %#x:%#x", vid, pid, fvid, fpid)
			err = errors.New(message)
			return true, fileData, err
		}

		return true, fileData[0 : fileSize-DFU_SUFFIX_LENGTH], nil
	}

	return false, fileData, nil
}

type MemoryLayout struct {
	start     int
	end       int
	totalSize int
	blockSize int
}

func parseMemoryLayout(str string) (layout MemoryLayout, err error) {
	layout = MemoryLayout{}
	sections := strings.Split(str, "/")
	if len(sections) != 3 {
		return layout, fmt.Errorf("[DFU] expected dfu string to contain three data sections separated by /, got '%s'", str)
	}

	if strings.IndexRune(str, '@') != 0 {
		return layout, fmt.Errorf("[DFU] expected dfu string to start with @, got '%s'", str)
	}

	start, err := strconv.ParseInt(sections[1], 0, 32)
	if err != nil {
		return layout, fmt.Errorf("[DFU] failed to extract start address '%s'", err)
	}
	layout.start = int(start)

	re := regexp.MustCompile(`([0-9]+)\s*\*\s*([0-9]+)\s?([ BKM])\s*([abcdefg])\s*,?\s*`)
	sectorLayout := re.FindStringSubmatch(str)

	if len(sectorLayout) < 4 {
		return layout, fmt.Errorf("[DFU] failed to extract sector layout '%s'", err)
	}

	sectorCount, err := strconv.ParseInt(sectorLayout[1], 0, 16)
	if err != nil {
		return layout, fmt.Errorf("[DFU] failed to extract sector count: '%s'", err)
	}

	sectorSize, err := strconv.ParseInt(sectorLayout[2], 0, 8)
	if err != nil {
		return layout, fmt.Errorf("[DFU] failed to extract sector size '%s'", err)
	}
	var multiplier int

	switch sectorLayout[3] {
	case " ":
	case "B":
		multiplier = 1
	case "K":
		multiplier = 1024
	case "M":
		multiplier = 1048576
	default:
		return layout, fmt.Errorf("[DFU] failed to extract sector multiplier from '%s'", str)
	}

	layout.blockSize = int(sectorSize) * multiplier
	layout.totalSize = layout.blockSize * int(sectorCount)
	layout.end = layout.start + layout.totalSize - 1

	return layout, nil

}

func (d *USBDevice) DFUFlash(firmwarePath string, cb func(message FlashCallback)) error {

	d.cb = cb
	d.flashStatus = dfuStatus{}

	d.cb(FlashCallback{Type: Log, Message: "[DFU] flashing process"})

	fileData, err := os.ReadFile(firmwarePath)
	if err != nil {
		return fmt.Errorf("Error while opening firmware: %s", err)
	}

	hasSuffix, firmwareData, err := d.extractSuffix(fileData)
	if err != nil {
		return err
	}

	if hasSuffix {
		d.cb(FlashCallback{Type: Log, Message: "[DFU] found a valid dfu suffix"})
	} else {
		d.cb(FlashCallback{Type: Log, Message: "[DFU] No DFU suffix found"})
	}

	claimed := d.Handle.Usb_claim()

	if !claimed {
		return errors.New("[DFU] couldn't claim the usb device")
	}

	defer d.Handle.Usb_close()

	/*
		dfu_str := d.Handle.Get_dfu_string(1)

		if dfu_str == "" {
			return errors.New("[DFU] couldn't read memory layout")
		}

		d.cb(FlashCallback{Type: Log, Message: fmt.Sprintf("[DFU] parsing memory layout from descriptor '%s'", dfu_str)})
		memory, err := parseMemoryLayout(dfu_str)
		if err != nil {
			return err
		}
	*/

	blockSize := 2048
	startAddress := STM32_START_ADDRESS

	config := d.Handle.Usb_set_configuration(1)
	if config > 0 {
		return fmt.Errorf("[DFU] couldn't set usb configuration %d", config)
	}

	err = d.dfuGetStatus()
	if err != nil {
		return err
	}

	if d.flashStatus.bStatus == DFU_ERROR {
		d.cb(FlashCallback{Type: Log, Message: "[DFU] clearing status"})
		err = d.dfuClearStatus()
		if err != nil {
			return err
		}

		err = d.dfuGetStatus()
		if err != nil {
			return err
		}

		if d.flashStatus.bStatus == DFU_ERROR {
			return errors.New("[DFU] Failed clearing status, try to unplug / plug the device and retry flashing")
		}

	}

	d.cb(FlashCallback{Type: Log, Message: "[DFU] status cleared"})

	d.cb(FlashCallback{Type: Log, Message: "[DFU] erasing flash"})
	err = d.dfuCommand(DFU_ERASE_SECTOR, 0)
	if err != nil {
		return err
	}

	d.cb(FlashCallback{Type: Log, Message: "[DFU] flash erased"})

	fileSize := len(firmwareData)
	d.cb(FlashCallback{Type: Log, Message: fmt.Sprintf("[DFU] ready to flash %d bytes", fileSize)})

	for page := 0; page < fileSize; page += blockSize {
		addr := startAddress + page
		chunckSize := blockSize

		if page+chunckSize > fileSize {
			chunckSize = fileSize - page
		}

		d.cb(FlashCallback{Type: Log, Message: fmt.Sprintf("[DFU] flashing memory at block 0x%X", addr)})

		err = d.dfuCommand(DFU_SET_ADDRESS, addr)
		if err != nil {
			return err
		}

		block_buffer := firmwareData[page : page+chunckSize]
		err = d.dfuDownload(block_buffer)
		if err != nil {
			return err
		}

		d.cb(FlashCallback{Type: Log, Message: fmt.Sprintf("[DFU] wrote %d bytes", chunckSize)})
		d.cb(FlashCallback{Type: Progress, Sent: page + chunckSize, Total: fileSize})
	}

	time.Sleep(1 * time.Second)

	err = d.dfuCommand(DFU_SET_ADDRESS, startAddress)

	if err != nil {
		return err
	}

	cb(FlashCallback{Type: Log, Message: "[DFU] memory pointer set to flash beginning"})

	d.dfuReboot()
	cb(FlashCallback{Type: Log, Message: "[DFU] reboot command sent"})
	return nil
}
