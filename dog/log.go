package dog

import (
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
    "from": "2022-09-10T11:48:36+01:00",
    "to": "2022-09-12T12:48:36+01:00"
  },
  "sort": "timestamp",
  "page": {
    "limit": 5
  }
}`
	payloadString := fmt.Sprintf(payload, query)
	jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payloadString))
	fmt.Println(jsonString)
}
