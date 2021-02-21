package influxv2

import (
	"context"
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

	return &Client {
		client: &client,
		options: &options,
	}, nil
}

func (influxClient *Client) Send(thing Influxable) {
	writeAPI := (*influxClient.client).WriteAPIBlocking(influxClient.options.Organization, influxClient.options.Bucket)

	p := influxdb2.NewPoint(thing.Category(), thing.Tags(), thing.Fields(), time.Now())
	writeAPI.WritePoint(context.Background(), p)
	(*influxClient.client).Close()
}