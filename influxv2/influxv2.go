package influxv2

import (
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
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
	client, err := influxdb2.NewClient(options.Server, options.Token)
	defer client.Close()

	if err != nil {
		return nil, err
	}

	return &Client {
		client: &client,
		options: &options,
	}, nil
}

func (influxv2Client *Client) Send(thing Influxable) error {
	writeAPI, err := (influxv2Client *Client).WriteAPI(options.Organization, options.Bucket)

	if err != nil {
		return err
	}

	p, err := influxdb2.NewPoint(thing.Category(), thing.Tags(), thing.Fields())

	if err != nil {
		return err
	}

	writeAPI.WritePoint(p)
	writeAPI.Flush()
}