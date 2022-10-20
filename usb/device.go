package usb

import (
	"fmt"
	"strings"
	"unsafe"
)

const LIBUSB_SUCCESS = 0x00
const REPORT_ID byte = 0x00
const ORYX_USAGE_PAGE = 0xFF60

const CMD_GET_FW_VERSION byte = 0x00
const CMD_INIT_PAIRING byte = 0x01
const CMD_VALIDATE_PAIRING byte = 0x02

const EVT_GET_FIRMWARE_VERSION byte = 0x00
const EVT_LAYER_CHANGE byte = 0x05
const EVT_KEY_PRESSED byte = 0x06
const EVT_KEY_RELEASED byte = 0x07

type USBDevice struct {
	Bootloader     bool                  `json:"bootloader"`
	Fingerprint    int                   `json:"fingerprint"`
	FirmwareFormat DeviceFirmware_format `json:"firmwareFormat"`
	FriendlyName   string                `json:"friendlyName"`
	Handle         Device
	Model          string `json:"model"`
	PortNumber     int
	Protocol       DeviceFlash_protocol
	cb             func(message FlashCallback)
	flashStatus    dfuStatus
}

func (d *USBDevice) Open() bool {
	connected := d.Handle.Hid_open(ORYX_USAGE_PAGE)
	if connected {
		cb := NewDirectorHIDPacketHandler(d)
		d.Handle.SetPacket_handler(cb)
		go func() {
			d.Handle.Hid_listen()
		}()
	} else {
		return false
	}
	fwVersionPacket := make([]byte, HID_PACKET_SIZE)

	fwVersionPacket[0] = REPORT_ID
	fwVersionPacket[1] = CMD_GET_FW_VERSION
	d.Handle.Send_hid_packet(&fwVersionPacket[0], HID_PACKET_SIZE)

	initPairingPacket := make([]byte, HID_PACKET_SIZE)

	initPairingPacket[0] = REPORT_ID
	initPairingPacket[1] = CMD_VALIDATE_PAIRING
	d.Handle.Send_hid_packet(&initPairingPacket[0], HID_PACKET_SIZE)
	return connected
}

func (d *USBDevice) Control(bmRequestType, bRequest uint8, wValue, wIndex uint16, data []byte) (buf []byte, err error) {

	var res TransferStatus
	size := len(data)

	if data == nil {
		res = d.Handle.Usb_transfer(bmRequestType, bRequest, wValue, wIndex, nil, 0, 0)

	} else {
		res = d.Handle.Usb_transfer(bmRequestType, bRequest, wValue, wIndex, &data[0], uint16(size), 0)
	}

	status_code := res.GetStatus_code()
	if status_code != LIBUSB_SUCCESS {
		status_error := res.GetStatus_error()
		return nil, fmt.Errorf("USB Transfer error: %s", status_error)
	}

	buffer := res.GetBuf()
	buffer_slice := unsafe.Slice(buffer, size)

	return buffer_slice, nil
}

func (d *USBDevice) Info() string {
	return fmt.Sprintf("Friendly name: '%s' | Fingerprint: %d | Port Number: %d | Bootloader: %t", d.FriendlyName, d.Fingerprint, d.PortNumber, d.Bootloader)
}

func parseIncomingPacket(packet *int8) (cmd byte, params []byte) {
	data := unsafe.Slice(packet, HID_PACKET_SIZE)
	cmd = byte(data[0])
	var i int
	for i = 0; i < HID_PACKET_SIZE; i++ { // exclude the first byte since it's the command
		if data[i] == -2 { // end of packet, we can now infer the packet params length using i
			break
		}
	}

	params = make([]byte, i-1)

	for j := 1; j < i; j++ {
		params[j-1] = byte(data[j])
	}

	return cmd, params
}

// TODO pass a callback to make stuff generic
func (d *USBDevice) HandleIncomingPacket(packet *int8) {
	cmd, params := parseIncomingPacket(packet)

	//fmt.Println("Received hid command", cmd)

	switch cmd {
	case EVT_GET_FIRMWARE_VERSION:
		v := string(params)

		version := strings.Split(v, "/")

		if len(version) == 2 {
			//uiEvent.Emit("firmwareVersion", &FirmwareVersionEvent{LayoutId: version[0], RevisionId: version[1]})
		}
		break
	case EVT_LAYER_CHANGE:
		//layerNum := params[0]
		//uiEvent.Emit("layerChanged", &LayerChangeEvent{LayerNum: layerNum})
		break
	case EVT_KEY_PRESSED:
		//uiEvent.Emit("keyPressed", &KeyPressEvent{Pressed: true, Column: params[0], Row: params[1]})
		break
	case EVT_KEY_RELEASED:
		//uiEvent.Emit("keyPressed", &KeyPressEvent{Pressed: false, Column: params[0], Row: params[1]})
		break
	}
}
