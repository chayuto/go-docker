package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/tarm/serial"
)

func main() {
	token := "GH0rNP89jL0aaU9ImXJ5Ko87VziUeSMwm0fsgR75cJGke9732oe5pW2bqzwQSKoLIBz59Mfdxf4c90-iq_Ou9A==" // os.Getenv("INFLUXDB_TOKEN")
	url := "https://ap-southeast-2-1.aws.cloud2.influxdata.com"
	client := influxdb2.NewClient(url, token)

	org := "nonesecure@gmail.com"
	bucket := "test1"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	c := &serial.Config{Name: "COM4", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(s)
		reply, _ := reader.ReadBytes('\x0a')
		fmt.Println(reply)

		readStr := string(reply)
		fmt.Printf("GOT: %v\n", readStr)

		value, err := strconv.ParseFloat(strings.TrimSpace(readStr), 32)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Val:: %v\n", value)

		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"Humidity": value,
		}
		point := write.NewPoint("measurement2", tags, fields, time.Now())

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Written %v\n", value)
		}
	}

}
