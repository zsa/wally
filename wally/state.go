package wally

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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

func NewState(step int8, filePath string) State {

	s := State{Step: step}
	s.AppVersion = GetAppVersion()

	if filePath != "" && runtime.GOOS != "darwin" {
		_, err := ioutil.ReadFile(filePath)
		if err != nil {
			message := fmt.Sprintf("Error while opening firmware: %s", err)
			s.Log("error", message)
		} else {
			extension := filepath.Ext(filePath)
			if extension == ".bin" {
				s.Device = Device{Model: 0, Bus: 0, Port: 0}

			} else if extension == ".hex" {
				s.Device = Device{Model: 1, Bus: 0, Port: 0}
			} else {
				message := fmt.Sprintf("File extension %s is not supported", extension)
				s.Log("error", message)
				return s
			}
			s.SelectFirmware(filePath)
		}
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

func (s *State) SelectFirmwareWithData(data string) {
	fileName := fmt.Sprintf("_wally_%d", time.Now().Unix())
	filePath := filepath.Join(os.TempDir(), fileName)
	dataStr := strings.Split(data, " ")
	var dataInt []int8
	buf := new(bytes.Buffer)
	for _, b := range dataStr {
		i, _ := strconv.ParseInt(b, 10, 8)
		dataInt = append(dataInt, int8(i))
	}
	err := binary.Write(buf, binary.LittleEndian, dataInt)
	err = ioutil.WriteFile(filePath, buf.Bytes(), 0644)
	if err != nil {
		message := fmt.Sprintf("Error while creating the temporary firmware file: %s", err)
		s.Log("error", message)
	} else {
		s.FirmwarePath = filePath
		s.Step = 3
	}
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
