package usb

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/marcinbor85/gohex"
)

const (
	HALFKAY_REPORT_ID      = 0x00
	HALFKAY_USAGE_PAGE     = 65436
	ERGODOX_MEM_SIZE       = 32256
	ERGODOX_SECTOR_SIZE    = 128
	HALFKAY_PACKET_RETRIES = 5
)

func (d *USBDevice) SendWithRetries(buf []byte, silent bool) (err error) {
	for retries := 0; retries < HALFKAY_PACKET_RETRIES; retries++ {
		_, _err := d.Control(0x21, 9, 0x0200, 0, buf)
		if _err != nil {
			time.Sleep(500 * time.Millisecond)
			if !silent {
				d.cb(FlashCallback{Type: Log, Message: fmt.Sprintf("[HALFKAY] sending packet failed: %s - retrying", err)})
			}
		} else {
			return nil
		}
	}
	return fmt.Errorf("error sending hid packet after %d retries", HALFKAY_PACKET_RETRIES)
}

func (d *USBDevice) HALFKAYFlash(firmwarePath string, cb func(message FlashCallback)) error {
	d.cb = cb
	d.cb(FlashCallback{Type: Log, Message: "[HALFKAY] flashing process"})
	file, err := os.Open(firmwarePath)
	if err != nil {
		return fmt.Errorf("Error while opening firmware: %s", err)
	}
	defer file.Close()

	firmwareData := gohex.NewMemory()
	d.cb(FlashCallback{Type: Log, Message: "[HALFKAY] parsing intel hex file"})
	err = firmwareData.ParseIntelHex(file)

	if err != nil {
		return fmt.Errorf("Error while parsing firmware: %s", err)
	}

	d.cb(FlashCallback{Type: Log, Message: "[HALFKAY] intel hex file parsed"})

	claimed := d.Handle.Usb_claim()
	if !claimed {
		return errors.New("[HALFKAY] couldn't claim the usb device")
	}
	defer d.Handle.Usb_close()

	config := d.Handle.Usb_set_configuration(1)
	if config > 0 {
		return fmt.Errorf("[HALFKAY] couldn't set usb configuration %d", config)
	}

	var addr uint32
	for addr = 0; addr < ERGODOX_MEM_SIZE; addr += ERGODOX_SECTOR_SIZE {
		// Prepare and write a firmware block
		// https://www.pjrc.com/teensy/halfkay_protocol.html
		buf := make([]byte, ERGODOX_SECTOR_SIZE+2)
		buf[0] = byte(addr & 255)
		buf[1] = byte((addr >> 8) & 255)
		sector := firmwareData.ToBinary(addr, ERGODOX_SECTOR_SIZE, 0xFF)
		for index := range sector {
			buf[index+2] = sector[index]
		}
		d.cb(FlashCallback{Type: Log, Message: fmt.Sprintf("[HALFKAY] writing to address 0x%X", addr)})

		err := d.SendWithRetries(buf, false)

		if err != nil {
			return err
		}

		// set a longer timeout when writing the first block
		if addr == 0 {
			time.Sleep(3 * time.Second)
		} else {
			time.Sleep(100 * time.Millisecond)
		}
		d.cb(FlashCallback{Type: Progress, Sent: int(addr), Total: ERGODOX_MEM_SIZE})
	}

	// send reboot packet
	buf := make([]byte, ERGODOX_SECTOR_SIZE+2)
	buf[0] = byte(0xFF)
	buf[1] = byte(0xFF)
	err = d.SendWithRetries(buf, true)

	if err != nil {
		return err
	}

	return nil
}
