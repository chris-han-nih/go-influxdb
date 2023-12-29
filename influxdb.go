package main

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"os"
	"time"
)

type InfluxDBClient struct {
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI
}

func NewInfluxDBClient() (*InfluxDBClient, error) {
	url := os.Getenv("INFLUXDB_URL")
	if url == "" {
		return nil, fmt.Errorf("INFLUXDB_URL is empty")
	}
	token := os.Getenv("INFLUXDB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("INFLUXDB_TOKEN is empty")
	}
	org := os.Getenv("INFLUXDB_ORG")
	if org == "" {
		return nil, fmt.Errorf("INFLUXDB_ORG is empty")
	}
	bucket := os.Getenv("INFLUXDB_BUCKET")
	if bucket == "" {
		return nil, fmt.Errorf("INFLUXDB_BUCKET is empty")
	}
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	queryAPI := client.QueryAPI(org)

	return &InfluxDBClient{
		client:   client,
		writeAPI: writeAPI,
		queryAPI: queryAPI,
	}, nil
}

func (c *InfluxDBClient) WriteData(
	measurement string,
	tags map[string]string,
	fields map[string]interface{},
) error {
	p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	return c.writeAPI.WritePoint(context.Background(), p)
}

func (c *InfluxDBClient) ReadData(
	ctx context.Context,
	query string,
) error {
	result, err := c.queryAPI.Query(ctx, query)
	if err != nil {
		return err
	}

	for result.Next() {
		fmt.Printf("값: %v, 시간: %v\n", result.Record().Value(), result.Record().Time())
	}

	return result.Err()
}

func (c *InfluxDBClient) Close() {
	c.client.Close()
}
