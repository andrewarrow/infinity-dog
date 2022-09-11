package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/network"
)

func Logs(query string) {
	payload := `{
  "filter": {
    "query": "%s",
    "indexes": [
      "main"
    ],
    "from": "2022-09-10T00:00:00+01:00",
    "to": "2022-09-12T23:59:59+01:00"
  },
  "sort": "timestamp",
  "page": {
    "limit": 50
  }
}`
	payloadString := fmt.Sprintf(payload, query)
	jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payloadString))

	var logResponse LogResponse
	json.Unmarshal([]byte(jsonString), &logResponse)

	for _, d := range logResponse.Data {
		fmt.Println(d.Attributes.Service)
	}
}
