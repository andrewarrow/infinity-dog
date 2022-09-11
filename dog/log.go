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
    "from": "2022-09-10T11:48:36+01:00",
    "to": "2022-09-12T12:48:36+01:00"
  },
  "sort": "timestamp",
  "page": {
    "limit": 5
  }
}`
	jsonString := network.DoPost("/api/v2/logs/events/search", []byte(payload))
	fmt.Println(jsonString)
}
