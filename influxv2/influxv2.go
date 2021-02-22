package influxv2

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb-client-go"
)

type Options struct {
	Bucket string
	Organization string
	Server string
	Token string
}

type Influxable interface {
	Tags() map[string]string
	Fields() map[string]interface{}
	Category() string
}

type Client struct {
	client *influxdb2.Client
	options *Options
}

func Connect(options Options) (*Client, error) {
	client := influxdb2.NewClient(options.Server, options.Token)
	defer client.Close()

	if client != nil {
		return &Client {
			client: &client,
			options: &options,
		}, nil
	}
	return nil, fmt.Errorf("could not connect to %v", options.Server)
}

func (influxClient *Client) Send(thing Influxable) {
	writeAPI := (*influxClient.client).WriteAPI(influxClient.options.Organization, influxClient.options.Bucket)
	
	// write errors if any
	errorChannel := writeAPI.Errors()
	go func() {
		for err := range errorChannel {
			og.Fatalln("error writing to influx", err.Error())
		}
	}()

	// write data
	p := influxdb2.NewPoint(thing.Category(), thing.Tags(), thing.Fields(), time.Now())
	writeAPI.WritePoint(p)
	writeAPI.Flush()
	(*influxClient.client).Close()
}