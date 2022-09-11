package dog

import (
	"fmt"
	"infinity-dog/network"
)

func Logs() {
	jsonString := network.DoGet("/api/v2/logs/events/search")
	fmt.Println(jsonString)
}
