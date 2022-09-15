package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/files"
	"infinity-dog/network"
	"strings"
	"time"
)

func golangTimeToDogTime(s string) string {
	dateString := s[0:10]
	timeString := s[11:19]
	return dateString + "T" + timeString
}

func Logs(hours int, query string) {

	utc, _ := time.LoadLocation("UTC")
	utcNow := time.Now().In(utc)
	utcString := fmt.Sprintf("%v", utcNow.Add(time.Hour*time.Duration(hours*-1)))
	from := golangTimeToDogTime(utcString)
	utcString = fmt.Sprintf("%v", utcNow.Add(time.Second))
	to := golangTimeToDogTime(utcString)

	cursor := ""
	buff := []string{}
	startTime := time.Now().Unix()
	hits := 0
	for {
		fmt.Println(from, to, cursor)
		payloadString := makePayload(query, from, to, cursor)
		// 300 requests per hour (aka 5 per minute)
		jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payloadString))
		hits++
		if hits == 5 {
			for {
				delta := time.Now().Unix() - startTime
				if delta > 60 {
					break
				}
				fmt.Println("at 5", delta)
				time.Sleep(time.Second * 1)
			}
			startTime = time.Now().Unix()
			hits = 0
		}

		buff = append(buff, jsonString)

		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)

		/*
			now := time.Now().Unix()
			for _, d := range logResponse.Data {
				delta := now - d.Attributes.Timestamp.Unix()
				tsFloat := float64(delta) / 60.0
				fmt.Printf("%.2f %s\n", tsFloat, d.Attributes.Service)
				fmt.Printf("%s\n", d.Attributes.Message)
				fmt.Printf("%s\n", d.Attributes.SubAttributes.Msg)
				fmt.Printf("%s\n", d.Attributes.SubAttributes.Exception)
			}*/

		cursor = logResponse.Meta.Page.After

		if cursor == "" {
			break
		}
	}
	files.SaveFile("t.json", strings.Join(buff, "\n"))
}

func makePayload(query, from, to, cursor string) string {
	payload := `{
  "filter": {
    "query": "%s",
    "indexes": [
      "main"
    ],
		"from": "%s+01:00",
    "to": "%s+01:00"
  },
  "sort": "timestamp",
  "page": {
	  "cursor": %s,
    "limit": 1000
  }
}`
	cursorString := "null"
	if cursor != "" {
		cursorString = fmt.Sprintf(`"%s"`, cursor)
	}
	return fmt.Sprintf(payload, query, from, to, cursorString)
}
