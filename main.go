package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	bucket := "nih"
	org := "nih"
	token := "gaPVLznc9DDJHe4wkyUj9h2Yfq5Mtsjf7bFiYAKYmmo615OWKL7sM6K0N6rN6mTduzppd4ao4IsHrpBc39prZA=="
	url := "http://localhost:8086"

	os.Setenv("INFLUXDB_URL", url)
	os.Setenv("INFLUXDB_TOKEN", token)
	os.Setenv("INFLUXDB_ORG", org)
	os.Setenv("INFLUXDB_BUCKET", bucket)

	client, err := NewInfluxDBClient()
	if err != nil {
		fmt.Printf("Error creating InfluxDB client: %s\n", err)
		return
	}
	defer client.Close()

	for {
		avgValue := rand.Intn(100000-1) + 100
		maxValue := rand.Intn(70000-2) * 100
		if err = client.WriteData("kale", map[string]string{"unit": "temperature"}, map[string]interface{}{"avg": avgValue, "max": maxValue}); err != nil {
			fmt.Printf("Error writing data: %s\n", err)
		}

		avgValue = rand.Intn(100000-1) + 100
		maxValue = rand.Intn(70000-2) * 100
		if err = client.WriteData("harvey", map[string]string{"unit": "temperature"}, map[string]interface{}{"avg": avgValue, "max": maxValue}); err != nil {
			fmt.Printf("Error writing data: %s\n", err)
		}

		avgValue = rand.Intn(100000-1) + 100
		maxValue = rand.Intn(70000-2) * 100
		if err = client.WriteData("bruce", map[string]string{"unit": "temperature"}, map[string]interface{}{"avg": avgValue, "max": maxValue}); err != nil {
			fmt.Printf("Error writing data: %s\n", err)
		}
		fmt.Printf("write point %v %v\n", avgValue, maxValue)
		time.Sleep(10 * time.Millisecond)

		//query := `
		//		from(bucket: "nih")
		//		  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)
		//		  |> filter(fn: (r) => r["_measurement"] == "bruce")
		//		  |> filter(fn: (r) => r["_field"] == "avg" or r["_field"] == "max")
		//		  |> filter(fn: (r) => r["unit"] == "temperature")
		//		  |> aggregateWindow(every: v.windowPeriod, fn: mean, createEmpty: false)
		//		  |> yield(name: "mean")`
		//if err = client.ReadData(context.Background(), query); err != nil {
		//	fmt.Printf("Error reading data: %s\n", err)
		//}
	}
	//client := influxdb2.NewClient(url, token)
	//writeAPI := client.WriteAPI(org, bucket)
	//for {
	//	avgValue := rand.Intn(100000-1) + 100
	//	maxValue := rand.Intn(70000-2) * 100
	//	fmt.Printf("write point %v %v\n", avgValue, maxValue)
	//	p := influxdb2.NewPoint("kale",
	//		map[string]string{"unit": "temperature"},
	//		map[string]interface{}{"avg": avgValue, "max": maxValue},
	//		time.Now())
	//	writeAPI.WritePoint(p)
	//	avgValue = rand.Intn(100000-1) + 100
	//	maxValue = rand.Intn(70000-2) + 100
	//	p1 := influxdb2.NewPoint("harvey",
	//		map[string]string{"unit": "temperature"},
	//		map[string]interface{}{"avg": avgValue, "max": maxValue},
	//		time.Now())
	//	writeAPI.WritePoint(p1)
	//
	//	avgValue = rand.Intn(100000-1) + 100
	//	maxValue = rand.Intn(70000-2) * 100
	//	p2 := influxdb2.NewPoint("bruce",
	//		map[string]string{"unit": "temperature"},
	//		map[string]interface{}{"avg": avgValue, "max": maxValue},
	//		time.Now())
	//	writeAPI.WritePoint(p2)
	//	time.Sleep(10 * time.Millisecond)
	//}
	//client.Close()
}
