package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type configStruct struct {
	Version  int     `yaml:"version"`
	Push     string  `yaml:"push"`
	Interval string  `yaml:"interval"`
	Nxapi    []Nxapi `yaml:"nxapi"`
}

// Nxapi
type Nxapi struct {
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	Host     []string `yaml:"host"`
	Port     int      `yaml:"port"`
	Protocol string   `yaml:"protocol"`
}

var version = ""

func stateSwitch(s string) (state int) {
	switch s {
	case "unknown":
		state = -1
	case "down":
		state = 0
	case "up":
		state = 1
	case "link-up":
		state = 2
	}
	return
}

func printError(err error, str ...interface{}) {
	if err == nil {
		return
	}
	log.Printf("Error: %v", err)
	log.Println(str...)
}

func readConfig(config_file string) (config configStruct, err error) {
	var yamlFile []byte

	// Load yamlFile
	yamlFile, err = ioutil.ReadFile(config_file)
	if err != nil {
		return
	}

	// Parse yamlFile
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return
	}

	//err = yaml.Unmarshal([]byte(data), &conf)
	for i, qryConf := range config.Nxapi {
		// Set some defaults
		if qryConf.Port == 0 {
			config.Nxapi[i].Port = 443
		}
		if qryConf.Protocol == "" {
			config.Nxapi[i].Protocol = "https"
		}
	}
	return
}
