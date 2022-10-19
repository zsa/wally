package usb

import "errors"

type FlashCallbackType int8

const (
	Log FlashCallbackType = iota
	Progress
)

type FlashCallback struct {
	Type    FlashCallbackType
	Message string
	Sent    int
	Total   int
}

func (d *USBDevice) Flash(firmwarePath string, cb func(message FlashCallback)) error {
	switch d.Protocol {
	case DeviceDFU:
		return d.DFUFlash(firmwarePath, cb)
	case DeviceHALFKAY:
		return d.HALFKAYFlash(firmwarePath, cb)
	default:
		return errors.New("unknown flash protocol")
	}
}
