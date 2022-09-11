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

	cursor := ""
	for {
		payloadString := makePayload(query, from, to, cursor)
		jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payloadString))

		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)

		now := time.Now().Unix()
		for _, d := range logResponse.Data {
			delta := now - d.Attributes.Timestamp.Unix()
			tsFloat := float64(delta) / 60.0
			fmt.Printf("%.2f %s\n", tsFloat, d.Attributes.Service)
		}

		cursor = logResponse.Meta.Page.After

		if cursor == "" {
			break
		}
	}
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
    "limit": 500
  }
}`
	cursorString := "null"
	if cursor != "" {
		cursorString = fmt.Sprintf(`"%s"`, cursor)
	}
	return fmt.Sprintf(payload, query, from, to, cursorString)
}
