package main

import (
	"context"
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {
	token := "GH0rNP89jL0aaU9ImXJ5Ko87VziUeSMwm0fsgR75cJGke9732oe5pW2bqzwQSKoLIBz59Mfdxf4c90-iq_Ou9A==" // os.Getenv("INFLUXDB_TOKEN")
	url := "https://ap-southeast-2-1.aws.cloud2.influxdata.com"
	client := influxdb2.NewClient(url, token)

	org := "nonesecure@gmail.com"
	bucket := "test1"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 0; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Written %d\n", value)
		}
	}

}
