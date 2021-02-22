# NUT Forwarder for InfluxDB

<a href="https://hub.docker.com/r/jstauffer/nut-forwarder-influxdb">
  <img alt="jstauffer/nut-forwarder-influxdb Docker Pulls" src="https://img.shields.io/docker/pulls/jstauffer/nut-forwarder-influxdb">
</a>
<a href="https://github.com/jknyght9/nut-forwarder-influxdb/blob/master/LICENSE">
  <img alt="AGPL-3.0 License" src="https://img.shields.io/github/license/jknyght9/nut-forwarder-influxdb">
</a>

> Fork of the [albinodrought/nut-forwarder-influxdb](https://github.com/AlbinoDrought/nut-forwarder-influxdb) repo. This addresses connectivity to InfluxDB version 2. Please use original project if using InfluxDB v1.

Forward some of your [Network UPS Tools (NUT)](https://networkupstools.org/index.html) data to InfluxDB. I used the NUT server running on a Synology NAS connected to a Tripp Lite UPS.

## Environment Variables

- `INFLUX_SERVER`: URL to InfluxDB server including scheme and port, defaults to `http://localhost:8086`
- `INFLUX_BUCKET`: Bucket to save data to, defaults to `ups`
- `INFLUX_ORGANIZATION`: InfluxDB organization
- `INFLUX_TOKEN`: InfluxDB user token in base64 encoding
- `NUT_HOST`: hostname or IP address of your NUT server, defaults to `localhost`
- `NUT_USERNAME`: NUT username, defaults to empty (no auth)
- `NUT_PASSWORD`: NUT password, defaults to empty (no auth)

## Building w/ GoLang

1. Ensure you have golang installed
2. May need to run command `go mod init main.go`

```
go get -d -v
go build
```

## Building w/ Docker

1. Run command `go mod init main.go`

```
go mod init main.go
docker build -t jknyght9/nut-forwarder-influxdb .
```

## Running

Program will attempt to use environment variables for runtime, (see Environment Variables). Based on your shell environment, set the variables and run the resulting binary.

#### Bash with remote InfluxDB and NUT servers

```bash
INFLUX_ORGANIZATION=myorg \
INFLUX_BUCKET=ups \
INFLUX_TOKEN=base64token== \
INFLUX_SERVER=http://remoteaddr:8086 \
NUT_HOST=remoteaddr \
./nut-forwarder-influxdb
```

#### Fish with remote InfluxDB and NUT servers

```fish
set -x INFLUX_ORGANIZATION myorg \
set -x INFLUX_BUCKET ups \
set -x INFLUX_TOKEN base64token== \
set -x INFLUX_SERVER http://remoteaddr:8086 \
set -x NUT_HOST remoteaddr \
./nut-forwarder-influxdb
```