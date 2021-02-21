package influxv2

import (
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

func (influxv2Client *Client) Send(thing Influxable) error {
	writeAPI, err := influxv2Client.WriteAPIBlocking(influxv2Client.options.Organization, influxv2Client.options.Bucket)

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