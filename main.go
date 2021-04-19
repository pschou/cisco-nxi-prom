package main

import (
	"bytes"
	"fmt"
	"github.com/pschou/go-cisco-nx-api/pkg/client"
	"github.com/pschou/go-params"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

type res struct {
	Result []byte
}

var queryStep, queryInterval time.Duration
var oneTime bool
var host_count int
var nextRun time.Time
var config configStruct
var config_file *string

func main() {
	// Define and parse arguments
	params.CommandLine.Title = "Cisco to Prometheus Collector, Version: " + version
	config_file = params.String("conf", "config.yml", "Config file to read from", "FILE")
	params.Parse()

	// Read in and parse the configs (see code below)
	readAndParseConfig()
	reconfig_sig := make(chan os.Signal, 1)
	signal.Notify(reconfig_sig, syscall.SIGHUP)
	go func() {
		for {
			<-reconfig_sig
			log.Println("Reloading config (via HUP signal)...")
			readAndParseConfig()
		}
	}()

	nextRun = time.Now()

	// Loop for query interval
	for {

		// Loop over config blocks
		for _, qryConf := range config.Nxapi {
			runtime.GC()

			// Loop over hosts in the config
			for _, host := range qryConf.Host {
				// Create buffer for output to send or to display
				var buf bytes.Buffer

				// Add delay for next query
				sleep_time := nextRun.Sub(time.Now())
				for sleep_time > queryStep {
					sleep_time -= queryStep
				}
				time.Sleep(sleep_time)
				nextRun = nextRun.Add(queryStep)

				go func() {
					log.Println("Querying host", host)
					cli := client.NewClient()
					cli.SetHost(host)
					cli.SetUsername(qryConf.User)

					// Alternate method of providing a password Environment variable
					password := qryConf.Password
					if len(password) > 0 && password[0] == '@' {
						dat, _ := ioutil.ReadFile(password[1:])
						password = string(dat)
					}
					if password == "" {
						password = os.Getenv("PASSWORD")
					}

					cli.SetPassword(password)
					cli.SetProtocol(qryConf.Protocol)
					cli.SetPort(qryConf.Port)

					results, err := cli.Configure([]string{
						"show version",          //done 2
						"show bgp session",      //done 2
						"show ip route",         //done 2
						"show ip arp",           //done 2
						"show interface status", //done 2
						"show interface quick",  //done 2
						"show isis adj detail",  //done 2
					})

					/* // Test data for development
					var err error
					results := make([]res, 7)
					results[0].Result, err = ioutil.ReadFile("/home/schou/git/go-cisco-nx-api/assets/requests/resp.result.show.version.json")
					results[1].Result, err = ioutil.ReadFile("/home/schou/git/go-cisco-nx-api/assets/requests/resp.result.show.bgp.sessions.json")
					results[2].Result, err = ioutil.ReadFile("/home/schou/git/go-cisco-nx-api/assets/requests/resp.result.show.ip.route.json")
					results[3].Result, err = ioutil.ReadFile("/home/schou/git/go-cisco-nx-api/assets/requests/resp.result.show.ip.arp.json")
					results[4].Result, err = ioutil.ReadFile("/home/schou/git/go-cisco-nx-api/assets/requests/resp.result.show.interface.status.json")
					results[5].Result, err = ioutil.ReadFile("/home/schou/git/go-cisco-nx-api/assets/requests/resp.result.show.interface.quick.json")
					results[6].Result, err = ioutil.ReadFile("/home/schou/git/go-cisco-nx-api/assets/requests/resp.result.show.isis.2.adj.det.json")
					*/

					if err != nil {
						log.Println("Error in call to API for host", host, "err", err)
						return
					}

					//for _, result := range results {
					//	fmt.Fprintf(&buf,"data %#v\n\n\n", string(result.Result)) // DEBUG
					//}

					ver_resp, err := client.NewVersionResultFromBytes(results[0].Result)
					printRespErr(err, "ver", results[0].Result)

					bgp_resp, err := client.NewBGPSessionsResultFromBytes(results[1].Result)
					printRespErr(err, "bgp", results[1].Result)

					iprt_resp, err := client.NewIpRouteResultFromBytes(results[2].Result)
					printRespErr(err, "iproute", results[2].Result)

					iparp_resp, err := client.NewIpArpResultFromBytes(results[3].Result)
					printRespErr(err, "iparp", results[3].Result)

					stat_resp, err := client.NewInterfaceStatusResultFromBytes(results[4].Result)
					printRespErr(err, "stat", results[4].Result)

					quick_resp, err := client.NewInterfaceQuickResultFromBytes(results[5].Result)
					printRespErr(err, "quick", results[5].Result)

					isis_resp, err := client.NewIsisAdjDetailResultFromBytes(results[6].Result)
					printRespErr(err, "isis", results[6].Result)

					//
					// Parse Version blob into metrics
					//
					if ver_resp != nil {
						fmt.Fprintf(&buf, "cisco_info{biosVer=%q,sysVer=%q,boardID=%q,chassisID=%q} 1\n",
							ver_resp.Body.BiosVerStr, ver_resp.Body.KickstartVerStr, ver_resp.Body.ProcBoardID, ver_resp.Body.ChassisID)

						const longForm = "Mon Jan 2 15:04:05 2006"
						t, _ := time.Parse(longForm, ver_resp.Body.RrCTime)

						fmt.Fprintf(&buf, "cisco_reset_time{rr_reason=%q,rr_service=%q,rr_sysVer=%q} %d\n",
							ver_resp.Body.RrReason, ver_resp.Body.RrService, ver_resp.Body.RrSysVer, t.Unix())

						fmt.Fprintf(&buf, "cisco_uptime_seconds %v\n",
							((ver_resp.Body.KernUptmDays*24+ver_resp.Body.KernUptmHrs)*60+
								ver_resp.Body.KernUptmMins)*60+ver_resp.Body.KernUptmSecs)

						switch ver_resp.Body.MemType {
						case "mB":
							fmt.Fprintf(&buf, "cisco_memory %v\n", 1024*1024*ver_resp.Body.Memory)
						case "kB":
							fmt.Fprintf(&buf, "cisco_memory %v\n", 1024*ver_resp.Body.Memory)
						default:
							fmt.Fprintf(&buf, "cisco_memory{memType=%q} %v\n", ver_resp.Body.MemType, ver_resp.Body.Memory)
						}
					}

					//
					// Parse BGP sessions into metrics
					//
					if bgp_resp != nil {
						bgp_slices := bgp_resp.Flat()
						for _, b := range bgp_slices {
							fmt.Fprintf(&buf, "cisco_bgp_lastflap_seconds{neighborID=%q,remoteAS=\"%d\",localAS=\"%d\",routerID=%q} %d\n",
								b.NeighborID, b.RemoteAS, b.LocalAS, b.RouterID, b.LastFlap/1e9)
							state := 0
							if b.State == "Established" {
								state = 1
							}
							fmt.Fprintf(&buf, "cisco_bgp_state{neighborID=%q,remoteAS=\"%d\",localAS=\"%d\",routerID=%q} %d\n",
								b.NeighborID, b.RemoteAS, b.LocalAS, b.RouterID, state)
							fmt.Fprintf(&buf, "cisco_bgp_conndrop_count{neighborID=%q,remoteAS=\"%d\",localAS=\"%d\",routerID=%q} %d\n",
								b.NeighborID, b.RemoteAS, b.LocalAS, b.RouterID, b.ConnectionsDropped)
						}
					}

					//
					// Parse IP route into metrics
					//
					if iprt_resp != nil {
						route_slices := iprt_resp.Flat()
						for _, r := range route_slices {
							fmt.Fprintf(&buf, "cisco_ip_route_uptime_seconds{clientName=%q,ifName=%q,ipPrefix=%q} %d\n",
								r.ClientName, r.IfName, r.IPPrefix, r.UpTime/1e9)
							fmt.Fprintf(&buf, "cisco_ip_route_mcast_hops{clientName=%q,ifName=%q,ipPrefix=%q} %d\n",
								r.ClientName, r.IfName, r.IPPrefix, r.MCastNHops)
							fmt.Fprintf(&buf, "cisco_ip_route_ucast_hops{clientName=%q,ifName=%q,ipPrefix=%q} %d\n",
								r.ClientName, r.IfName, r.IPPrefix, r.UCastNHops)
							fmt.Fprintf(&buf, "cisco_ip_route_pref{clientName=%q,ifName=%q,ipPrefix=%q} %d\n",
								r.ClientName, r.IfName, r.IPPrefix, r.Pref)
							fmt.Fprintf(&buf, "cisco_ip_route_metric{clientName=%q,ifName=%q,ipPrefix=%q} %d\n",
								r.ClientName, r.IfName, r.IPPrefix, r.Metric)
						}
					}

					//
					// Parse IP ARP into metrics
					//
					if iparp_resp != nil {
						arp_slices := iparp_resp.Flat()
						for _, r := range arp_slices {
							fmt.Fprintf(&buf, "cisco_ip_arp{flags=%q,intfOut=%q,iPAddrOut=%q,mac=%q} %d\n",
								r.Flags, r.IntfOut, r.IPAddrOut, r.MAC, r.TimeStamp/1e9)
						}
					}

					//
					// Parse interface status into metrics
					//
					if stat_resp != nil {
						stat_slices := stat_resp.Flat()
						for _, r := range stat_slices {
							fmt.Fprintf(&buf, "cisco_interface_speed_bits{interface=%q,state=%q,vlan=%q,type=%q,autoSpeed=\"%v\"} %d\n",
								r.Interface, r.State, r.VLAN, r.Type, r.SpeedAuto, r.SpeedVal)
						}
					}

					//
					// Parse interface quick into metrics
					//
					if quick_resp != nil {
						quick_slices := quick_resp.Flat()
						for _, r := range quick_slices {
							lbl := fmt.Sprintf("interface=%q,mac=%q", r.Interface, r.EthHwAddr)
							fmt.Fprintf(&buf, "cisco_interface_info{%s,desc=%q,eth_autoneg=%q} 1\n",
								lbl, r.Desc, r.EthAutoNeg)

							fmt.Fprintf(&buf, "cisco_interface_state{%s} %d\n", lbl, stateSwitch(r.State))
							fmt.Fprintf(&buf, "cisco_interface_admin_state{%s} %d\n", lbl, stateSwitch(r.AdminState))

							for i, v := range r.EthBW {
								fmt.Fprintf(&buf, "cisco_interface_bw_bits{%s,stream=\"%d\"} %d\n", lbl, i, v)
							}

							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_in_pkts{%s} %d\n", lbl, r.VdcLvlInPkts)
							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_in_bytes{%s} %d\n", lbl, r.VdcLvlInBytes)
							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_in_ucast_pkts{%s} %d\n", lbl, r.VdcLvlInUCast)
							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_in_mcast_pkts{%s} %d\n", lbl, r.VdcLvlInMCast)
							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_in_bcast_pkts{%s} %d\n", lbl, r.VdcLvlInBCast)

							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_out_pkts{%s} %d\n", lbl, r.VdcLvlOutPkts)
							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_out_bytes{%s} %d\n", lbl, r.VdcLvlOutBytes)
							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_out_ucast_pkts{%s} %d\n", lbl, r.VdcLvlOutUCast)
							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_out_mcast_pkts{%s} %d\n", lbl, r.VdcLvlOutMCast)
							fmt.Fprintf(&buf, "cisco_interface_vdc_lvl_out_bcast_pkts{%s} %d\n", lbl, r.VdcLvlOutBCast)
						}
					}

					//
					// Parse ISIS status into metrics
					//
					if isis_resp != nil {
						isis_slices := isis_resp.Flat()
						for _, r := range isis_slices {
							fmt.Fprintf(&buf, "cisco_isis_adj_transitions{intfOut=%q,iPAddrOut=%q,iP6AddrOut=%q} %d\n",
								r.AdjIntfNameOut, r.AdjIpv4AddrOut, r.AdjIpv6AddrOut, r.AdjTransitionsOut)
						}
					}

					if config.Push == "" {
						// Print out the result
						fmt.Printf("metrics:\n%s", buf.String())
					} else {
						// Send the result to Prometheus Collector
						UploadToCollector(strings.TrimSuffix(config.Push, "/")+"/host/"+host, buf.Bytes())
					}
				}()
			}
		}
		if oneTime == true {
			break
		}
	}
}

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
