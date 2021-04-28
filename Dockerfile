ARG ARCH="amd64"
ARG OS="linux"
FROM scratch
LABEL description="Cisco NXI exporter for Prometheus, built by Paul Schou (paulschou.com)"

ADD ./LICENSE /LICENSE
ADD ./cisco-prom "/cisco-prom"
ENTRYPOINT  [ "/cisco-prom" ]
