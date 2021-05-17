package main

import (
	"log"
	"time"
)

// Handle the loading and parsing the timing for the config
func readAndParseConfig() {
	var err error

	// Read Config
	log.Println("Loading the configuration file.")
	config, err = readConfig(*config_file)
	printError(err, "Error reading or parsing config file:", *config_file, "error:", err)

	// Parse the interval
	if len(config.Interval) > 0 {
		queryInterval, err = time.ParseDuration(config.Interval)
		printError(err, "Parsing interval error:", err)
	} else {
		// One time shot deal
		oneTime = true
	}

	//Count the total number of hosts to query
	host_count = 0
	for _, qryConf := range config.Nxapi {
		for range qryConf.Host {
			host_count++
		}
	}
	log.Println("Found", host_count, "hosts in the config.")

	if host_count == 0 {
		log.Fatal("No hosts found, please add hosts for querying")
	}

	// Figure out the timing
	if !oneTime {
		queryStep = time.Duration(int64(queryInterval) / int64(host_count))
	}
}

// Print error to the screen
func printRespErr(err error, t string, dat []byte) {
	if err != nil {
		log.Println(t, "=", string(dat))
		log.Println("err=", err)
	}
}
