package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/files"
	"infinity-dog/network"
	"infinity-dog/util"
	"os"
	"strconv"
	"time"
)

func Device(deviceId string) {
	hoursAsInt, _ := strconv.Atoi(os.Getenv("DOG_BASE"))
	if hoursAsInt == 0 {
		hoursAsInt = 24
	}

	// Calculate last 10 minutes in PT
	pt, _ := time.LoadLocation("America/Los_Angeles")
	ptNow := time.Now().In(pt)
	ptFrom := ptNow.Add(-10 * time.Minute)
	
	// Display time range in PT format
	fmt.Printf("Time range: %s – %s PT\n", 
		ptFrom.Format("Jan 2, 3:04 pm"), 
		ptNow.Format("Jan 2, 3:04 pm"))

	// Then call Logs with the device-specific query for last 10 minutes
	query := fmt.Sprintf("service:aroma-broker @device_id:%s", deviceId)
	DeviceLogsMinutes(10, query)
}

func DeviceLogsMinutes(minutes int, query string) {
	utc, _ := time.LoadLocation("UTC")
	utcNow := time.Now().In(utc)
	// we seem to be off by about 1 hour
	utcNow = utcNow.Add(time.Minute * 55)

	// Go back exactly the specified number of minutes
	utcString := fmt.Sprintf("%v", utcNow.Add(time.Minute*time.Duration(minutes*-1)))
	from := golangTimeToDogTime(utcString)
	utcString = fmt.Sprintf("%v", utcNow.Add(time.Second))
	to := golangTimeToDogTime(utcString)

	cursor := ""
	startTime := time.Now().Unix()
	hits := 0
	for {
		fmt.Println(from, to, cursor)
		payloadString := makePayload(query, from, to, cursor)
		// 300 requests per hour (aka 5 per minute)
		jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payloadString))
		hits++
		if hits == 300 {
			for {
				delta := time.Now().Unix() - startTime
				if delta > 3600 {
					break
				}
				fmt.Println("at 300", delta)
				time.Sleep(time.Second * 1)
			}
			startTime = time.Now().Unix()
			hits = 0
		}

		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)

		// Display filtered results for device command
		for _, d := range logResponse.Data {
			if d.Attributes.SubAttributes.PayloadHex != "" {
				fmt.Printf("payload_hex: %s\n", d.Attributes.SubAttributes.PayloadHex)
				fmt.Printf("topic: %s\n", d.Attributes.SubAttributes.Topic)
				fmt.Printf("time: %s\n", d.Attributes.SubAttributes.Time)
				fmt.Println("---")
			}
		}

		cursor = logResponse.Meta.Page.After

		if cursor == "" {
			break
		}
	}
}

func LogsMinutes(minutes int, query string) {
	utc, _ := time.LoadLocation("UTC")
	utcNow := time.Now().In(utc)
	// we seem to be off by about 1 hour
	utcNow = utcNow.Add(time.Minute * 55)

	// Go back exactly the specified number of minutes
	utcString := fmt.Sprintf("%v", utcNow.Add(time.Minute*time.Duration(minutes*-1)))
	from := golangTimeToDogTime(utcString)
	utcString = fmt.Sprintf("%v", utcNow.Add(time.Second))
	to := golangTimeToDogTime(utcString)

	cursor := ""
	startTime := time.Now().Unix()
	hits := 0
	for {
		fmt.Println(from, to, cursor)
		payloadString := makePayload(query, from, to, cursor)
		// 300 requests per hour (aka 5 per minute)
		jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payloadString))
		hits++
		if hits == 300 {
			for {
				delta := time.Now().Unix() - startTime
				if delta > 3600 {
					break
				}
				fmt.Println("at 300", delta)
				time.Sleep(time.Second * 1)
			}
			startTime = time.Now().Unix()
			hits = 0
		}

		files.SaveFile(fmt.Sprintf("samples/%s.json", util.PseudoUuid()), jsonString)

		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)

		cursor = logResponse.Meta.Page.After

		if cursor == "" {
			break
		}
	}
}
