package state

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
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
	if currentVersion != u.Version {
		return true
	}
	return false
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

	fmt.Println(updateUrl)
	if res.StatusCode != 200 {
		return update, fmt.Errorf("failed to fetch update data from server, error code %d", res.StatusCode)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return update, err
	}

	err = json.Unmarshal(body, &update)
	fmt.Println(string(body[:]))

	return update, nil
}
