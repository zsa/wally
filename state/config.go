package state

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
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

func getFilePath() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}
	return path.Join(configPath, FILE_NAME+".yml"), nil
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

	filePath, err := getFilePath()

	if err != nil {
		return "", created, err
	}

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

type Configuration struct {
	firstrun bool

	UpdateCheck bool `yaml:"update_check"`
}

func (c *Configuration) SetUpdateCheck(val bool) {
	c.UpdateCheck = val
	err := c.saveConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Configuration) saveConfig() error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	filepath, err := getFilePath()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, bytes, 0644)
}

func (c *Configuration) readConfig() error {
	filepath, err := getFilePath()
	if err != nil {
		return err
	}
	f, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(f, c)
}

func NewConfiguration() Configuration {
	_, firstrun, _ := createConfigFile()
	if firstrun {
		return Configuration{firstrun: firstrun, UpdateCheck: false}
	} else {
		config := Configuration{firstrun: false, UpdateCheck: false}
		config.readConfig()
		return config
	}
}
