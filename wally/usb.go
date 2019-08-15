package wally

import (
	"fmt"
	"github.com/google/gousb"
)

type Device struct {
	Model int `json:"model"` // 0 - planck // 1 - ergodox
	Bus   int `json:"bus"`
	Port  int `json:"port"`
}

const (
	vendorID  int = 0xFEED
	planckID  int = 0x6060
	ergodoxID int = 0x1307

	dfuSuffixVendorID  int = 0x83
	dfuSuffixProductID int = 0x11
	dfuVendorID        int = 0x0483
	dfuProductID       int = 0xdf11

	halfKayVendorID  int = 0x16C0
	halfKayProductID int = 0x0478

	ergodoxMaxMemorySize = 0x100000
	ergodoxCodeSize      = 32256
	ergodoxBlockSize     = 128

	dfuSuffixLength    = 16
	planckBlockSize    = 2048
	planckStartAddress = 0x08000000
	setAddress         = 0
	eraseAddress       = 1
	eraseFlash         = 2
)

func ProbeDevices(s *State) []Device {
	devices := []Device{}
	ctx := gousb.NewContext()
	defer ctx.Close()
	s.Log("info", "Probing compatible usb devices")

	_, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		if desc.Vendor == gousb.ID(vendorID) {
			if desc.Product == gousb.ID(planckID) {
				devices = append(devices, Device{Model: 0, Bus: desc.Bus, Port: desc.Port})

			}
			if desc.Product == gousb.ID(ergodoxID) {
				devices = append(devices, Device{Model: 1, Bus: desc.Bus, Port: desc.Port})
			}
		}

		if desc.Vendor == gousb.ID(dfuVendorID) && desc.Product == gousb.ID(dfuProductID) {
			devices = append(devices, Device{Model: 0, Bus: desc.Bus, Port: desc.Port})
		}

		if desc.Vendor == gousb.ID(halfKayVendorID) && desc.Product == gousb.ID(halfKayProductID) {
			devices = append(devices, Device{Model: 1, Bus: desc.Bus, Port: desc.Port})
		}

		return false
	})

	if err != nil {
		message := fmt.Sprintf("OpenDevices: %s", err)
		s.Log("warning", message)
	}

	message := fmt.Sprintf("Found %d compatible device(s)", len(devices))
	s.Log("info", message)

	return devices
}
