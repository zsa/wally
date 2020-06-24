package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/wailsapp/wails"
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
	Total int `json:"total"` // total of firmware bytes to send
	Sent  int `json:"sent"`  // total of bytes sent
}

//State represents the global state of the application
type State struct {
	runtime       *wails.Runtime
	AppVersion    string        `json:"appVersion"`
	Device        Device        `json:"device"`        // The user selected usb device
	Devices       []Device      `json:"devices"`       // The list of usb devices connected
	Step          int8          `json:"step"`          // The current flashing process step. // 0 - Probing keyboard // 1 - Select keyboard // 2 - Select firmware file // 3 - Waiting for keyboard reset // 4 - Flashing // 5 - Complete
	FirmwarePath  string        `json:"firmwarePath"`  // The firmware absolute Path selected by the user
	FlashProgress FlashProgress `json:"flashProgress"` // The Flashing state progress
	Logs          []log         `json:"logs"`          // Log object
}

func NewState(step int8, filePath string) *State {
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
				return &s
			}
		}
	}
	return &s
}

func (s *State) WailsInit(runtime *wails.Runtime) error {
	s.runtime = runtime
	runtime.Events.On("wails:loaded", func(...interface{}) {
		s.emitUpdate()
		s.ProbeDevices()
	})
	return nil
}

func (s *State) emitUpdate() {
	s.runtime.Events.Emit("state_update", s)
}

func (s *State) Log(level string, message string) {
	now := time.Now()
	l := log{Timestamp: now.Unix(), Level: level, Message: message}
	s.Logs = append(s.Logs, l)
	s.emitUpdate()
}

func (s *State) ProbeDevices() {
	for len(s.Devices) == 0 {
		s.Devices = ProbeDevices(s)
		if len(s.Devices) > 1 {
			s.Step = 1
			s.emitUpdate()
		}
		if len(s.Devices) == 1 {
			s.Device = s.Devices[0]
			s.Step = 2
			s.emitUpdate()
		}
		time.Sleep(1 * time.Second)
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
	s.Step = 0
	s.ProbeDevices()
	s.Log("info", "Application state reset")
}

func (s *State) SelectDevice(model int, bus int, port int) {
	device := Device{Model: model, Bus: bus, Port: port}
	s.Device = device
	s.Step = 2
	s.emitUpdate()
}

func (s *State) SelectFirmware() {
	filter := ""
	if s.Device.Model == 1 {
		filter = "*.hex"
	} else {
		filter = "*.bin"
	}
	s.FirmwarePath = s.runtime.Dialog.SelectFile("Select a firmware file", filter)
	fmt.Println("Select")
	fmt.Println(s.FirmwarePath)

	if s.FirmwarePath != "" {
		s.Step = 3
		s.FlashFirmware()
		s.emitUpdate()
	}
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
		s.FlashFirmware()
		s.emitUpdate()
	}
}

func (s *State) Shutdown() {
	s.runtime.Window.Close()
}

func (s *State) FlashFirmware() {
	if s.Device.Model == 1 {
		s.Log("info", "Starting Teensy Flash")
		go TeensyFlash(s)
	} else {
		s.Log("info", "Starting DFU Flash")
		go DFUFlash(s)
	}
}
