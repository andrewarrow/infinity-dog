package dog

import (
	"fmt"
	"infinity-dog/network"
)

func Logs() {
	payload := `{
  "filter": {
    "query": "datadog-agent",
    "indexes": [
      "main"
    ],
    "from": "2020-09-17T11:48:36+01:00",
    "to": "2020-09-17T12:48:36+01:00"
  },
  "sort": "timestamp",
  "page": {
    "limit": 5
  }
}`
	jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payload))
	fmt.Println(jsonString)
}
