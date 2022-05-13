package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	TwitchChannel  string
	PressDuration  int32 // Specified because GIU is weird about ints
	HoldDirections bool
	AllowSOCD      bool

	ButtonTriggers struct {
		A string
		B string

		X string
		Y string
		Z string

		L     string
		R     string
		START string

		UP    string
		DOWN  string
		LEFT  string
		RIGHT string

		CUP    string
		CDOWN  string
		CLEFT  string
		CRIGHT string
	}
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
