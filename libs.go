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

func PrintErr(err error, str ...interface{}) {
	if err == nil {
		return
	}
	log.Printf("Error: %v", err)
	log.Println(str...)
}

func readConfig(config_file string) (config configStruct) {
	// Load yamlFile
	yamlFile, err := ioutil.ReadFile(config_file)
	if err != nil {
		log.Fatal("Cannot read config file", config_file, err)
	}

	// Parse yamlFile
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal("Invalid config file format", err)
	}

	// Set a value for interval if it is not defined
	if config.Interval == "" {
		config.Interval = "1m"
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
