package wally

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	return string(b[1 : len(b)-1])
}

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

type Step int8

const (
	Probing        Step = iota // probing keyboard
	SelectKeyboard             // select keyboard
	FirmwareFile               // select firmware file
	Waiting                    // waiting for keyboard reset
	Flashing                   // flashing
	Complete                   // complete
)

//State represents the global state of the application
type State struct {
	runtime       *wails.Runtime
	AppVersion    string        `json:"appVersion"`
	Device        Device        `json:"device"`        // The user selected usb device
	Devices       []Device      `json:"devices"`       // The list of usb devices connected
	Step          Step          `json:"step"`          // The current flashing process step
	FirmwarePath  string        `json:"firmwarePath"`  // The firmware absolute Path selected by the user
	FlashProgress FlashProgress `json:"flashProgress"` // The Flashing state progress
	Logs          []log         `json:"logs"`          // Log object
}

func NewState(step Step, filePath string) *State {
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
	l := log{Timestamp: now.Unix(), Level: level, Message: jsonEscape(message)}
	s.Logs = append(s.Logs, l)
	s.emitUpdate()
}

func (s *State) ProbeDevices() {
	for len(s.Devices) == 0 {
		s.Devices = ProbeDevices(s)
		if len(s.Devices) > 1 {
			s.Step = SelectKeyboard
			s.emitUpdate()
		}
		if len(s.Devices) == 1 {
			s.Device = s.Devices[0]
			s.Step = FirmwareFile
			s.emitUpdate()
		}
		time.Sleep(1 * time.Second)
	}
}

func (s *State) CompleteFlash() {
	s.Step = Complete
}

func (s *State) ResetState() {
	s.Device = Device{}
	s.Devices = []Device{}
	s.FirmwarePath = ""
	s.FlashProgress = FlashProgress{}
	s.Step = Probing
	s.ProbeDevices()
	s.Log("info", "Application state reset")
}

func (s *State) SelectDevice(model int, bus int, port int) {
	device := Device{Model: model, Bus: bus, Port: port}
	s.Device = device
	s.Step = FirmwareFile
	s.emitUpdate()
}

func (s *State) SelectFirmware() {
	filter := ""
	if s.Device.Model == 1 {
		filter = "*.hex"
	} else {
		filter = "*.bin"
	}
	filePath := s.runtime.Dialog.SelectFile("Select a firmware file", filter)

	s.FirmwarePath = jsonEscape(filePath)

	if s.FirmwarePath != "" {
		s.Step = Waiting
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
		s.FirmwarePath = jsonEscape(filePath)
		s.Step = Flashing
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
