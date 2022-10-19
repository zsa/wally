package state

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"time"

	wails "github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/zsa/wally/usb"
)

type Step int8

const (
	Probing          Step = iota // probing keyboard
	KeybordSelect                // select keyboard
	FirmwareSelect               // select firmware file
	KeyboardReset                // waiting for keyboard reset
	FirmwareFlashing             // flashing
	FlashComplete                // complete
	FatalError
	WallyUpdate
	WallyUpdateComplete
)

type State struct {
	Devices        map[int]usb.USBDevice `json:"devices"`
	FirmwarePath   *string
	Logs           []log          `json:"logs"`
	SelectedDevice *usb.USBDevice `json:"selectedDevice"`
	Step           Step
	version        string
	ctx            context.Context
	enumerator     usb.Enumerator
	config         Configuration
	updatePath     string
}

func (s *State) Init(ctx context.Context) {
	s.ctx = ctx
	s.version = GetAppVersion()
	s.config = NewConfiguration()
}

func (s *State) GetAppVersion() string {
	return s.version
}

func (s *State) StateStart() {
	InitUIEventEmitter(s.ctx)
	enumerator := usb.NewEnumerator()
	cb := usb.NewDirectorEventHandler(s)
	enumerator.SetEventObject(cb)

	s.enumerator = enumerator

	if s.config.firstrun {
		res, err := wails.MessageDialog(s.ctx, wails.MessageDialogOptions{
			Buttons:       []string{"No", "Yes"},
			Type:          wails.QuestionDialog,
			Title:         "Wally updates",
			DefaultButton: "Yes",
			Message:       "Would you like Wally to check for updates on startup?",
		})
		if err != nil {
			res = "No"
		}
		fmt.Println(res)

	} else {
		update, err := checkForUpdate()
		if err != nil {
			s.Log("warning", fmt.Sprintf("failed to check for update: %s", err))
		}
		if update.required(s.version) {
			res, err := wails.MessageDialog(s.ctx, wails.MessageDialogOptions{
				Buttons:       []string{"No", "Yes"},
				Type:          wails.QuestionDialog,
				Title:         fmt.Sprintf("Version %s of Wally is available", update.Version),
				DefaultButton: "Yes",
				Message:       "Would you like to update now?",
			})
			if err != nil {
				res = "No"
			}

			if res == "Yes" {
				destination, err := s.DownloadUpdate(update)
				if err != nil {
					s.Log("fatal", err.Error())
				}
				s.updatePath = destination
				s.SetStep(WallyUpdateComplete)
			}
		}

	}

	go func() {
		s.Log("info", "UI started, listening to usb events")
		enumerator.Listen()
	}()

}

func (s *State) Log(level string, message string) {
	now := time.Now()
	l := log{Timestamp: now.Unix(), Level: level, Message: message}
	s.Logs = append(s.Logs, l)
	uiEvent.Emit("log", &LogEvent{Log: l})
}

func (s *State) SetStep(step Step) {
	s.Step = step
	uiEvent.Emit("stepChanged", &StepChangeEvent{Step: step})
}

func (s *State) SelectDevice(fingerprint int) {
	if device, ok := s.Devices[fingerprint]; ok {
		s.SelectedDevice = &device
		uiEvent.Emit("deviceSelected", &DeviceSelectedEvent{Device: *s.SelectedDevice})
	}
}

func (s *State) SelectFirmware() {

	if s.SelectedDevice == nil {
		//LOG err
		return
	}

	pattern := ""

	if s.SelectedDevice.FirmwareFormat == usb.DeviceBIN {
		pattern = "*.bin"
	}

	if s.SelectedDevice.FirmwareFormat == usb.DeviceHEX {
		pattern = "*.hex"
	}

	if pattern == "" {
		//Log err
		return
	}

	selection, err := wails.OpenFileDialog(s.ctx, wails.OpenDialogOptions{
		Title: "Select firmware file",
		Filters: []wails.FileFilter{{
			DisplayName: "Firmware file (" + pattern + ")",
			Pattern:     pattern,
		}},
	})

	if err != nil {
		//Log err
		return
	}

	if selection != "" {
		s.FirmwarePath = &selection
		// If the selected device is a bootloader start the flashing process
		// else jump to the Keyboard Reset screen
		if s.SelectedDevice.Bootloader {
			s.StartFlashing()
		} else {
			s.SetStep(KeyboardReset)
		}
	}
}

func (s *State) StartFlashing() {
	s.SetStep(FirmwareFlashing)
	err := s.SelectedDevice.Flash(*s.FirmwarePath, func(message usb.FlashCallback) {
		if message.Type == usb.Log {
			s.Log("info", message.Message)
		}
		if message.Type == usb.Progress {
			uiEvent.Emit("flashProgress", &ProgressEvent{Current: message.Sent, Total: message.Total})
		}
	})
	if err != nil {
		s.Log("fatal", err.Error())
		s.SetStep(FatalError)
	} else {
		s.Log("info", "flash complete")
		s.SetStep(FlashComplete)
	}
}

func (s *State) HandleUSBConnectionEvent(connect bool, dev usb.Device) {
	if s.Devices == nil {
		s.Devices = make(map[int]usb.USBDevice)
	}

	fingerprint := dev.GetFingerprint()
	name := dev.GetFriendly_name()
	bootloader := dev.GetBootloader()
	firmwareFormat := dev.GetFile_format()
	protocol := dev.GetProtocol()
	portNumber := dev.GetPort_number()

	if connect {
		device := usb.USBDevice{FriendlyName: name, Fingerprint: fingerprint, PortNumber: portNumber, Bootloader: bootloader, FirmwareFormat: firmwareFormat, Protocol: protocol, Handle: dev}
		s.Devices[fingerprint] = device
		uiEvent.Emit("deviceConnected", &DeviceConnectionEvent{Device: device})
		s.Log("info", "New device detected:")
		s.Log("info", fmt.Sprintf("'%s' | pointer: %d", name, fingerprint))
		s.Log("info", fmt.Sprintf("port: %d | bootloader: %t", portNumber, bootloader))

		// Trigger the flashing process if:
		// - The current step is KeyboardReset
		// - The connected device is on the same port number as the selected device
		// - The connected device is a bootloader
		if bootloader && s.SelectedDevice != nil && s.SelectedDevice.PortNumber == portNumber && s.Step == KeyboardReset {
			s.SelectedDevice = &device
			time.Sleep(1 * time.Second)
			s.SelectDevice(fingerprint)
			s.StartFlashing()
		}
	} else {
		uiEvent.Emit("deviceDisconnected", &DeviceDisconnectionEvent{Fingerprint: fingerprint})
	}
}

func (s *State) DownloadUpdate(update Update) (string, error) {
	s.SetStep(WallyUpdate)
	file := "wally-v" + update.Version
	switch runtime.GOOS {
	case "darwin":
		file += ".dmg"
	case "windows":
		file += "-installer.exe"
	}
	destination := path.Join(os.TempDir(), file)

	out, err := os.Create(destination)
	if err != nil {
		return "", fmt.Errorf("unable to create destination file, please try to download the update manually from: %s", update.URL)
	}

	defer out.Close()

	headRes, err := http.Head(update.URL)
	if err != nil {
		return "", fmt.Errorf("unable to contact download server, please try to download the update manually from: %s", update.URL)
	}
	defer headRes.Body.Close()

	fileSize, err := strconv.Atoi(headRes.Header.Get("Content-Length"))
	if err != nil {
		return "", fmt.Errorf("invalid response from download server, please try to download the update manually from: %s", update.URL)
	}

	done := make(chan int64)

	go func() {
		stop := false
		f, err := os.Open(destination)
		if err != nil {
			return
		}
		defer f.Close()

		for {
			select {
			case <-done:
				stop = true
			default:

				fi, err := f.Stat()
				if err != nil {
					stop = true
					break
				}

				size := fi.Size()

				if size == 0 {
					size = 1
				}

				uiEvent.Emit("updateProgress", &ProgressEvent{Current: int(size), Total: fileSize})
			}
			if stop {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	res, err := http.Get(update.URL)
	if err != nil {
		done <- 0
		return "", err
	}
	defer res.Body.Close()

	n, err := io.Copy(out, res.Body)

	if err != nil {
		done <- 0
		return "", fmt.Errorf("error while transfering download to local file, please try to download the update manually from: %s", update.URL)
	}

	done <- n
	uiEvent.Emit("updateProgress", &ProgressEvent{Current: fileSize, Total: fileSize})
	return destination, nil
}

func (s *State) InstallUpdate() {

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("hdiutil", "attach", s.updatePath)
	case "windows":
		cmd = exec.Command(s.updatePath)
	}

	err := cmd.Run()
	if err != nil {
		s.Log("fatal", err.Error())
		return
	}

	wails.Quit(s.ctx)
}

func (s *State) Open(fingerprint int) bool {
	dev := s.Devices[fingerprint]
	return dev.Open()
}

func (s *State) Reset() {
	s.SelectedDevice = nil
	s.FirmwarePath = nil
	uiEvent.Emit("reset", nil)
	s.SetStep(Probing)
}

func (s *State) Quit() {
	s.Teardown()
	wails.Quit(s.ctx)
}

func (s *State) Teardown() {
	//usb.DeleteEnumerator(s.enumerator)
}
