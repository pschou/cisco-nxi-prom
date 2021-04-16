package main

import (
	"bytes"
	"fmt"
	"github.com/pschou/go-cisco-nx-api/pkg/client"
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
			buf.Reset()

			cli := client.NewClient()
			cli.SetHost(host)
			cli.SetUsername(qryConf.User)
			cli.SetPassword(qryConf.Password)
			cli.SetProtocol(qryConf.Protocol)
			cli.SetPort(qryConf.Port)

			results, err := cli.GetGeneric([]string{
				"show version",          //done
				"show bgp session",      //done
				"show ip route",         //done
				"show ip arp",           //done
				"show interface status", //done
				"show interface quick",  //almost, have brief not quick
				"show isis adj detail",  //almost, have isis 2 adj detail
			})

			if err != nil {
				fmt.Println("Error in call to API for host", host, "err", err)
				continue
			}

			//for _, result := range results {
			//	fmt.Printf("data %#v\n", string(result.Result)) // DEBUG
			//}

			bgp_resp, err := client.NewBGPSessionFromBytes(results[1].Result)
			if err != nil {
				fmt.Println("Error in parsing BGP sessions for host", host, "err", err)
				continue
			}

			bgp_slices := bgp_resp.Flat()
			for _, bgp_slice := range bgp_slices {
				fmt.Printf("neighbor === %#v\n", bgp_slice)
			}
			// Print metric
			//buf.WriteString(fmt.Sprintf("CiscoProm_up{version=%q,host=%q} 1\n",
			//	SysInfo.Bios.Version, host))

			if config.Push == "" {
				// Print out the result
				fmt.Printf("metrics:\n%s", buf.String())
			} else {
				//UploadToCollector(config.Push, buf.String())
			}

		}
	}

}
