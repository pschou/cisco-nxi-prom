# cisco-nxi-prom

This is a Cisco NXI exporter for Prometheus.  It expects an endpoint for posting metrics, a Prometheus Collector.

As this is a directed tool, there is little here in terms of configuration on what is in and what isn't.

## Syntax
```
$ ./cisco-prom -h
Cisco to Prometheus Collector, Version: 0.1...

Usage: cisco-prom [option]
Option:
  --conf FILE  Config file to read from  (Default: "config.yml")
```

## Example config

This is a configuration file for doing a single shot grab of metrics without sending them to a Prometheus Collector:
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

This is a configuration file for grabbing metrics at a regular interval and sending them to a Prometheus Collector:
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
  password: @password.txt
```

Note: By putting @ in front of a file name one can reference a password from the contents of a file.  Otherwise, one may specify a file by leaving password blank and setting the environment variable, like PASSWORD=pass.

Fields used here are:
- port - The listening port on the network device
- protocol - The protocol used on the port on the network device (usually http/https)
- host - List of hosts to query for the metric
- user/password - Credentials to use for the scraping

To trigger a reload of a config file without restarting the server, use a `pkill -HUP cisco-prom`.
