package main

type config struct {
	InfluxServer   string `env:"INFLUX_SERVER" envDefault:"http://localhost:8086"`
	
	//v1
	// InfluxDatabase string `env:"INFLUX_DATABASE" envDefault:"ups"`
	// InfluxUsername string `env:"INFLUX_USERNAME"`
	// InfluxPassword string `env:"INFLUX_PASSWORD"`

	//v2
	// InfluxVersion2	bool `env:"INFLUX_VERSION2" envDefault:"true"`
	InfluxBucket   string `env:"INFLUX_BUCKET" envDefault:"ups"`
	InfluxOrganization   string `env:"INFLUX_ORGANIZATION"`
	InfluxToken   string `env:"INFLUX_TOKEN"`

	// NUT
	NUTHost     string `env:"NUT_HOST" envDefault:"localhost"`
	NUTUsername string `env:"NUT_USERNAME"`
	NUTPassword string `env:"NUT_PASSWORD"`
}
