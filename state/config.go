package state

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

const (
	FILE_NAME = "config"
)

func getConfigPath() (string, error) {
	homeDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configPath := path.Join(homeDir, "wally")
	return configPath, nil
}

func createConfigFile() (string, bool, error) {
	configPath, err := getConfigPath()
	created := false
	if err != nil {
		return "", created, err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.MkdirAll(configPath, 0700)
		if err != nil {
			return "", created, err
		}
	}

	filePath := path.Join(configPath, FILE_NAME+".yml")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.Create(filePath)
		if err != nil {
			return "", created, err
		}
		defer f.Close()
		created = true
	}
	return filePath, created, err
}

type Parameters struct {
	updates bool
}

type Configuration struct {
	new      bool
	firstrun bool
	params   Parameters
}

func (c *Configuration) PrompToCheckUpdates() bool {
	return c.new
}

func NewConfiguration() Configuration {
	_, firstrun, _ := createConfigFile()
	viper.SetConfigName(FILE_NAME)
	configPath, err := getConfigPath()
	if err != nil {
		viper.AddConfigPath(configPath)
	}
	return Configuration{firstrun: firstrun}
}
