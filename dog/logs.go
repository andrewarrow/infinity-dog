package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/network"
	"time"
)

func golangTimeToDogTime(s string) string {
	dateString := s[0:10]
	timeString := s[11:19]
	return dateString + "T" + timeString
}

func Logs(query string) {

	utc, _ := time.LoadLocation("UTC")
	utcNow := time.Now().In(utc)
	utcString := fmt.Sprintf("%v", utcNow.Add(time.Second*-60))
	from := golangTimeToDogTime(utcString)
	utcString = fmt.Sprintf("%v", utcNow.Add(time.Second))
	to := golangTimeToDogTime(utcString)

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
	  "cursor": null,
    "limit": 500
  }
}`
	payloadString := fmt.Sprintf(payload, query, from, to)
	jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payloadString))

	var logResponse LogResponse
	json.Unmarshal([]byte(jsonString), &logResponse)

	for _, d := range logResponse.Data {
		fmt.Println(d.Attributes.Timestamp, d.Attributes.Service)
	}
}
