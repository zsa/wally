package wally

import (
	"fmt"
	"github.com/google/gousb"
	"github.com/marcinbor85/gohex"
	"os"
	"time"
)

// TeensyFlash: Flashes Teensy boards.
// It opens the firmware file at the provided path, checks it's integrity, wait for the keyboard to be in Flash mode, flashes it and reboots the board.
func TeensyFlash(s *State) {
	file, err := os.Open(s.FirmwarePath)
	if err != nil {
		message := fmt.Sprintf("Error while opening firmware: %s", err)
		s.Log("error", message)
		return
	}
	defer file.Close()

	s.FlashProgress.Total = ergodoxCodeSize

	firmware := gohex.NewMemory()
	err = firmware.ParseIntelHex(file)
	if err != nil {
		message := fmt.Sprintf("Error while parsing firmware: %s", err)
		s.Log("error", message)
		return
	}

	ctx := gousb.NewContext()
	defer ctx.Close()
	var dev *gousb.Device

	// Loop until a keyboard is ready to flash
	for {
		s.Log("info", "Waiting for a DFU capable device")
		// if the app is reset stop this goroutine and close the usb context
		if s.Step != 3 {
			s.Log("info", "App reset, interrupting the flashing process.")
			return
		}

		devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
			if desc.Vendor == gousb.ID(halfKayVendorID) && desc.Product == gousb.ID(halfKayProductID) {
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

	// Detach keyboard from the kernel
	dev.SetAutoDetach(true)

	// Claim usb device
	cfg, err := dev.Config(1)
	defer cfg.Close()
	if err != nil {
		message := fmt.Sprintf("Error while claiming the usb interface: %s", err)
		s.Log("error", message)
		return
	}

	s.Step = 4
	s.emitUpdate()

	// Loop on the firmware data and program
	var addr uint32
	for addr = 0; addr < ergodoxCodeSize; addr += ergodoxBlockSize {
		// set a longer timeout when writing the first block
		if addr == 0 {
			dev.ControlTimeout = 5 * time.Second
		} else {
			dev.ControlTimeout = 500 * time.Millisecond
		}
		// Prepare and write a firmware block
		// https://www.pjrc.com/teensy/halfkay_protocol.html
		buf := make([]byte, ergodoxBlockSize+2)
		buf[0] = byte(addr & 255)
		buf[1] = byte((addr >> 8) & 255)
		block := firmware.ToBinary(addr, ergodoxBlockSize, 255)
		for index := range block {
			buf[index+2] = block[index]
		}

		bytes, err := dev.Control(0x21, 9, 0x0200, 0, buf)
		if err != nil {
			message := fmt.Sprintf("Error while sending firmware bytes: %s", err)
			s.Log("error", message)
			return
		}

		message := fmt.Sprintf("Sent %d bytes out of %d", addr, ergodoxCodeSize)
		s.Log("info", message)
		s.FlashProgress.Sent += bytes
		s.emitUpdate()
	}

	time.Sleep(1 * time.Second)

	s.Log("info", "Sending the reboot command")
	buf := make([]byte, ergodoxBlockSize+2)
	buf[0] = byte(0xFF)
	buf[1] = byte(0xFF)
	buf[2] = byte(0xFF)
	_, err = dev.Control(0x21, 9, 0x0200, 0, buf)

	if err != nil {
		message := fmt.Sprintf("Error while rebooting device: %s", err)
		s.Log("error", message)
		return
	}

	s.Step = 5
	s.Log("info", "Flash complete")
}
