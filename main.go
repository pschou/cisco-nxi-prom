package main

import (
	"bytes"
	"fmt"
	"github.com/greenpau/go-cisco-nx-api/pkg/client"
	"github.com/pschou/go-params"
)

func main() {
	// Define and parse arguments
	params.CommandLine.Title = "Cisco to Prometheus Collector, Version: " + version
	var config_file = params.String("conf", "config.yml", "Config file to read from", "FILE")
	params.Parse()

	// Read Config
	config := readConfig(*config_file)

	var buf bytes.Buffer

	// Loop over config blocks
	for _, qryConf := range config.Nxapi {
		// Loop over hosts in the config
		for _, host := range qryConf.Host {
			cli := client.NewClient()
			cli.SetHost(host)
			cli.SetPort(qryConf.Port)
			cli.SetProtocol(qryConf.Protocol)
			cli.SetUsername(qryConf.User)
			cli.SetPassword(qryConf.Password)

			// Get system information
			SysInfo, err := cli.GetSystemInfo()
			PrintErr(err)

			// Print metric
			buf.WriteString(fmt.Sprintf("CiscoProm_up{version=%q} 1\n", SysInfo.Bios.Version))

		}
	}

	if config.Push == "" {
		// Print out the result
		fmt.Printf("metrics:\n%s", buf)
	}
}
