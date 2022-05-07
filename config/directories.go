package config

import (
	"io/ioutil"
	"os"
	"path"
)

func configFileDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(configDir, "twitchToGamecube"), nil
}

func configFilePath() string {
	dir, _ := configFileDir()

	return path.Join(dir, "config.json")
}

func createBlankConfig() {
	dir, err := configFileDir()
	if err != nil {
		panic(err)
	}

	os.MkdirAll(dir, 0700)

	ioutil.WriteFile(configFilePath(), []byte("{}"), 0700)
}
