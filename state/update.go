package state

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	wails "github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	baseUrl = "https://configure.zsa.io/"
)

type Update struct {
	Version      string `json:"version"`
	URL          string `json:"url"`
	ReleaseNotes string `json:"releaseNotes"`
}

func (u *Update) required(currentVersion string) bool {
	return currentVersion != u.Version
}

func getUpdateUrl() string {
	filename := ""
	switch runtime.GOOS {
	case "darwin":
		filename = "wally-macos"
	case "linux":
		filename = "wally-linux"
	case "windows":
		filename = "wally-windows"
	}

	return fmt.Sprintf("%s%s.json", baseUrl, filename)
}

func checkForUpdate() (Update, error) {
	var update Update

	updateUrl := getUpdateUrl()
	res, err := http.Get(updateUrl)
	if err != nil {
		return update, err
	}

	if res.StatusCode != 200 {
		return update, fmt.Errorf("failed to fetch update data from server, error code %d", res.StatusCode)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return update, err
	}

	err = json.Unmarshal(body, &update)

	return update, err
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

func (s *State) PromptUpdates() {
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
		if res == "Yes" {
			s.config.SetUpdateCheck(true)
		}
	}
}

func (s *State) CheckUpdate() {
	if s.config.UpdateCheck {
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
					s.SetStep(FatalError)
				} else {
					s.updatePath = destination
					s.SetStep(WallyUpdateComplete)
				}
			}
		}
	}
}
