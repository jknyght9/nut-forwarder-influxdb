package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/jknyght9/nut-forwarder-influxdb/influxv2"
	// "github.com/AlbinoDrought/nut-forwarder-influxdb/influx"
	"github.com/caarlos0/env"
	nut "github.com/robbiet480/go.nut"
)

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("error parsing config", err)
	}

	client, err := nut.Connect(cfg.NUTHost)
	if err != nil {
		log.Fatalln("error connecting to NUT", err)
	}
	defer client.Disconnect()

	if cfg.NUTUsername != "" || cfg.NUTPassword != "" {
		authenticated, err := client.Authenticate(cfg.NUTUsername, cfg.NUTPassword)
		if err != nil {
			log.Fatalln("error during NUT authentication", err)
		}

		if !authenticated {
			log.Fatalln("NUT authentication failed")
		}
	}

	version, err := client.GetVersion()
	if err != nil {
		log.Fatalln("error getting NUT version", err)
	}

	log.Println("Connected to NUT!", version)

	influxv2Client, err := influx.Connect(influxv2.Options{
		Bucket:			cfg.InfluxBucket,
		Organization:	cfg.InfluxOrganization,
		Server:			cfg.InfluxServer,
		Token:			cfg.InfluxToken
	})
	if err != nil {
		log.Fatalln("error connecting to influx", err)
	}

	var upsList []nut.UPS
	ticker := time.NewTicker(time.Minute)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		upsList, err = client.GetUPSList()
		if err != nil {
			log.Fatalln("error getting UPS list", err)
		}

		for _, ups := range upsList {
			mapped := mapUPS(&ups)
			influxable := influxableUPS(mapped)

			err = influxv2Client.Send(influxable)
			if err != nil {
				log.Fatalln("error sending to influx", err)
			}
		}

		select {
		case <-ticker.C:
			// continue sending stats
			continue
		case <-interrupt:
			// stop processing, return from entire `main()`
			return
		}
	}
}
