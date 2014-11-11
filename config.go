package main

import (
	"encoding/json"
	"io/ioutil"
)

var (
	ConfigFile string
)

type ClientConfig struct {
	Journald struct {
		Sleep        int64
		Match        []string
		TriggerWords []string
	}

	Notifications struct {
		Email struct {
			Host     string
			Port     int
			Username string
			Password string
			To       []string
		}
	}
}

func (cfg *ClientConfig) Validate() (err error) {

	return
}

func GetCFG(f string) (cfg ClientConfig, err error) {

	b, err := ioutil.ReadFile(f)
	if err != nil {

		return
	}

	err = json.Unmarshal(b, &cfg)
	if err != nil {

		return
	}

	err = cfg.Validate()
	if err != nil {

		return
	}

	return
}
