package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	TwitchChannel  string
	SerialPortName string
}

func LoadConfig() Config {
	data, err := ioutil.ReadFile(configFilePath())
	if err != nil {
		createBlankConfig()
		return LoadConfig()
	}

	result := Config{}

	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	return result
}

func (c *Config) Save() {
	jsonData, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(configFilePath(), jsonData, 0700)
}
