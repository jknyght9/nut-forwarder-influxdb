package main

type config struct {
	InfluxServer   string `env:"INFLUX_SERVER" envDefault:"http://localhost:8086"`
	InfluxBucket   string `env:"INFLUX_BUCKET" envDefault:"ups"`
	InfluxOrganization   string `env:"INFLUX_ORGANIZATION"`
	InfluxToken   string `env:"INFLUX_TOKEN"`

	NUTHost     string `env:"NUT_HOST" envDefault:"localhost"`
	NUTUsername string `env:"NUT_USERNAME"`
	NUTPassword string `env:"NUT_PASSWORD"`
}
