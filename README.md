# cisco-nxi-prom

This is a Cisco NXI exporter for Prometheus.  It expects an endpoint for
posting metrics, a Prometheus Collector.  The main goal of this utility is to
be a lightweight metrics scraper capable of running in a very small memory
footprint.

As this is a simple / directed tool, there is little here in terms of
configuration.

This tool was built to work with the Prometheus Collector.  The offical source
and documentation can be found here: https://github.com/pschou/prom-collector

## Syntax
```
$ ./cisco-prom -h
Cisco to Prometheus Collector, Version: 0.1...

Usage: cisco-prom [option]
Option:
  --conf FILE  Config file to read from  (Default: "config.yml")
```

## Example config

This is a configuration file for doing a single shot grab of metrics without
sending them to a Prometheus Collector:
```
---
version: 1
nxapi:
- protocol: https
  host:
  - "host1"
  - "host2"
  port: 443
  user: myuser
  password: mypass
```

This is a configuration file for grabbing metrics at a regular interval and
sending them to a Prometheus Collector:
```
---
version: 1
push: http://localhost:9550/collector
interval: 5m
nxapi:
- protocol: https
  host:
  - "host1"
  - "host2"
  port: 443
  user: myuser
  password: "@password1.txt"
- host:
  - "host3"
  - "host4"
  user: secondSet
  password: "@password2.txt"
```

Note: By putting @ in front of a file name one can reference a password from
the contents of a file.  Otherwise, one may specify a file by leaving password
blank and setting the environment variable, like PASSWORD=pass.

Fields used here are:
- port - The listening port on the network device
- protocol - The protocol used on the port on the network device (usually http/https)
- host - List of hosts to query for the metric
- user/password - Credentials to use for the scraping

To trigger a reload of a config file without restarting the server, use a `pkill -HUP cisco-prom`.


# Example output
```
cisco_info{biosVer="08.32",sysVer="7.0(3)I7(4)",boardID="SAL2015NQ3H",chassisID="Nexus9000 C9508 (8 Slot) Chassis"} 1
cisco_reset_time{rr_reason="Reset Requested by CLI command reload",rr_service="",rr_sysVer="7.0(3)I7(4)"} 1527099972
cisco_uptime_seconds 18204
cisco_memory 16794398720
cisco_bgp_lastflap_seconds{neighborID="19.0.101.1",remoteAS="333",localAS="333",routerID="19.0.0.6"} 527611
cisco_bgp_state{neighborID="19.0.101.1",remoteAS="333",localAS="333",routerID="19.0.0.6"} 1
cisco_bgp_conndrop_count{neighborID="19.0.101.1",remoteAS="333",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_lastflap_seconds{neighborID="19.0.102.3",remoteAS="888",localAS="333",routerID="19.0.0.6"} 527611
cisco_bgp_state{neighborID="19.0.102.3",remoteAS="888",localAS="333",routerID="19.0.0.6"} 1
cisco_bgp_conndrop_count{neighborID="19.0.102.3",remoteAS="888",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_lastflap_seconds{neighborID="19.0.102.4",remoteAS="333",localAS="333",routerID="19.0.0.6"} 527611
cisco_bgp_state{neighborID="19.0.102.4",remoteAS="333",localAS="333",routerID="19.0.0.6"} 1
cisco_bgp_conndrop_count{neighborID="19.0.102.4",remoteAS="333",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_lastflap_seconds{neighborID="19.0.103.10",remoteAS="999",localAS="333",routerID="19.0.0.6"} 527611
cisco_bgp_state{neighborID="19.0.103.10",remoteAS="999",localAS="333",routerID="19.0.0.6"} 1
cisco_bgp_conndrop_count{neighborID="19.0.103.10",remoteAS="999",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_lastflap_seconds{neighborID="19.0.103.20",remoteAS="333",localAS="333",routerID="19.0.0.6"} 527608
cisco_bgp_state{neighborID="19.0.103.20",remoteAS="333",localAS="333",routerID="19.0.0.6"} 1
cisco_bgp_conndrop_count{neighborID="19.0.103.20",remoteAS="333",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_lastflap_seconds{neighborID="19.0.200.200",remoteAS="0",localAS="333",routerID="19.0.0.6"} 527625
cisco_bgp_state{neighborID="19.0.200.200",remoteAS="0",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_conndrop_count{neighborID="19.0.200.200",remoteAS="0",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_lastflap_seconds{neighborID="fec0::1002",remoteAS="333",localAS="333",routerID="19.0.0.6"} 527628
cisco_bgp_state{neighborID="fec0::1002",remoteAS="333",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_conndrop_count{neighborID="fec0::1002",remoteAS="333",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_lastflap_seconds{neighborID="fec0::2002",remoteAS="888",localAS="333",routerID="19.0.0.6"} 527628
cisco_bgp_state{neighborID="fec0::2002",remoteAS="888",localAS="333",routerID="19.0.0.6"} 0
cisco_bgp_conndrop_count{neighborID="fec0::2002",remoteAS="888",localAS="333",routerID="19.0.0.6"} 0
cisco_ip_route_uptime_seconds{clientName="static",ifName="Null0",ipPrefix="7.57.0.0/16"} 648125
cisco_ip_route_mcast_hops{clientName="static",ifName="Null0",ipPrefix="7.57.0.0/16"} 0
cisco_ip_route_ucast_hops{clientName="static",ifName="Null0",ipPrefix="7.57.0.0/16"} 1
cisco_ip_route_pref{clientName="static",ifName="Null0",ipPrefix="7.57.0.0/16"} 1
cisco_ip_route_metric{clientName="static",ifName="Null0",ipPrefix="7.57.0.0/16"} 0
cisco_ip_route_uptime_seconds{clientName="direct",ifName="Vlan253",ipPrefix="7.57.253.0/30"} 648034
cisco_ip_route_mcast_hops{clientName="direct",ifName="Vlan253",ipPrefix="7.57.253.0/30"} 0
```
