package wally

import (
	"fmt"

	"github.com/google/gousb"
)

// The Model type encodes which type of keyboard is being flashed. Three
// associated constants are defined for this type:
// Planck, ErgoDox, and Moonlander
type Model int

const (
	Planck Model = iota
	ErgoDox
	Moonlander
)

type Device struct {
	Model Model `json:"model"` // 0 - planck // 1 - ergodox // 2 - moonlander
	Bus   int   `json:"bus"`
	Port  int   `json:"port"`
}

const (
	vendorID1    int = 0xFEED
	vendorID2    int = 0x3297
	planckID     int = 0x6060
	ergodoxID    int = 0x1307
	moonlanderID int = 0x1969

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

var vendorIds = []int{0xfeed, 0x3297}                  // legacy - zsa's vendor id
var ergodoxIds = []int{0x1307, 0x4974, 0x4975, 0x4976} // legacy - standard - shine - glow
var planckIds = []int{0x6060, 0xc6ce, 0xc6cf}          // legacy - standard - glow
var moonlanderIds = []int{0x1969}                      // mk1

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ProbeDevices(s *State) []Device {
	devices := []Device{}
	ctx := gousb.NewContext()
	defer ctx.Close()
	s.Log("info", "Probing compatible usb devices")

	_, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		if contains(vendorIds, int(desc.Vendor)) {
			if contains(planckIds, int(desc.Product)) {
				devices = append(devices, Device{Model: 0, Bus: desc.Bus, Port: desc.Port})
			}

			if contains(ergodoxIds, int(desc.Product)) {
				devices = append(devices, Device{Model: 1, Bus: desc.Bus, Port: desc.Port})
			}

			if contains(moonlanderIds, int(desc.Product)) {
				devices = append(devices, Device{Model: 2, Bus: desc.Bus, Port: desc.Port})
			}
		}

		if desc.Vendor == gousb.ID(dfuVendorID) && desc.Product == gousb.ID(dfuProductID) {
			devices = append(devices, Device{Model: 3, Bus: desc.Bus, Port: desc.Port})
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
