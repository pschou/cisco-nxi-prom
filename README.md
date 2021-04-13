# cisco-nxi-prom

This is a Cisco NXI exporter for Prometheus.  It expects an endpoint for posting metrics, a Prometheus Collector.

## Syntax
```
$ ./cisco-prom -h
Cisco to Prometheus Collector, Version: 0.1...

Usage: cisco-prom [option]
Option:
  --conf FILE  Config file to read from  (Default: "config.yml")
```

## Example config
```
---
version: 1
push:
interval: 5m
nxapi:
- protocol: https
  host:
  - "host1"
  - "host2"
  port: 443
  user: myuser
  password: mypass
```
