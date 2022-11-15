package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var shared = *_Configuration

type _Configuration struct {
	Server struct {
		Port         int           `json:"port"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
	} `josn:"server"`

	Log struct {
		Verbose bool `json:"verbose"`
	} `json:"log"`
}

func init() {
	if shared != nil {
		return
	}

	basepath, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	bts, err := ioutil.ReadFile(filepath.Join(basepath, "conf", "config.json"))
	if err != nil {
		panic(err)
		return
	}

	shared = new(_Configuration)
	err = json.Unmarshal(bts, &shared)
	if err != nil {
		panic(err)
		return
	}
}

func Configuration() _Configuration {
	return *shared
}
