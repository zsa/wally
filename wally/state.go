package wally

import (
	"path/filepath"
	"runtime"
	"time"
)

type log struct {
	Timestamp int64  `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

//FlashProgress represents the current flashing state, it gets updated by the flashing methods.
type FlashProgress struct {
	Step  int `json:"step"`  // 0 - Probing keyboards // 1 - Flashing // 2 - Rebooting // 3 - Complete
	Total int `json:"total"` // total of firmware bytes to send
	Sent  int `json:"sent"`  // total of bytes sent
}

//State represents the global state of the application
type State struct {
	AppVersion    string        `json:"appVersion"`
	Device        Device        `json:"device"`        // The user selected usb device
	Devices       []Device      `json:"devices"`       // The list of usb devices connected
	Step          int8          `json:"step"`          // The current flashing process step. // 0 - Probing keyboard // 1 - Select keyboard // 2 - Select firmware file // 3 - Waiting for keybiard reset // 4 - Flashing // 5 - Complete
	FirmwarePath  string        `json:"firmwarePath"`  // The firmware absolute Path selected by the user
	FlashProgress FlashProgress `json:"flashProgress"` // The Flashing state progress
	Logs          []log         `json:"logs"`          // Log object
}

func NewState(step int8) State {
	s := State{Step: step}
	switch os := runtime.GOOS; os {
	case "darwin":
		s.AppVersion = "1.1.0"
	case "linux":
		s.AppVersion = "1.1.0"
	case "windows":
		s.AppVersion = "1.1.0"
	default:
		s.AppVersion = "1.1.0"
	}
	return s
}

func (s *State) Log(level string, message string) {
	now := time.Now()
	l := log{Timestamp: now.Unix(), Level: level, Message: message}
	s.Logs = append(s.Logs, l)
}

func (s *State) ProbeDevices() {
	s.Devices = ProbeDevices(s)
	if len(s.Devices) > 1 {
		s.Step = 1
	}
	if len(s.Devices) == 1 {
		s.Device = s.Devices[0]
		s.Step = 2
	}
}

func (s *State) PollFlashProgress() {
	state := s
	s = state
	if s.FlashProgress.Step == 3 {
		s.Step = 5
	}
}

func (s *State) CompleteFlash() {
	s.Step = 5
}

func (s *State) ResetState() {
	s.Device = Device{}
	s.Devices = []Device{}
	s.FirmwarePath = ""
	s.FlashProgress = FlashProgress{}
	s.Log("info", "Application state reset")
	s.Step = 0
}

func (s *State) SelectDevice(model int, bus int, port int) {
	device := Device{Model: model, Bus: bus, Port: port}
	s.Device = device
	s.Step = 2
}

func (s *State) SelectFirmware(path string) {
	s.FirmwarePath = path
	extension := filepath.Ext(path)
	if s.Device.Model == 0 && extension != ".bin" {
		return
	}

	if s.Device.Model == 1 && extension != ".hex" {
		return
	}

	s.Step = 3
}

func (s *State) FlashFirmware() {
	if s.Device.Model == 0 {

		s.Log("info", "Starting DFU Flash")
		go DFUFlash(s.FirmwarePath, s)
	}
	if s.Device.Model == 1 {
		s.Log("info", "Starting Teensy Flash")
		go TeensyFlash(s.FirmwarePath, s)
	}
}
