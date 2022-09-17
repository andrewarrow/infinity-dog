package dog

import (
	"fmt"
	"infinity-dog/database"
)

var servicesExceptions = []string{}
var servicesMessages = []string{}

func ServicesFromSql(sortString, service string) {
	serviceItems := database.ServicesByTotalBytes()
	for i, item := range serviceItems {
		fmt.Printf("%03d. %-60s %d\n", i+1, item.Name, item.TotalBytes)
	}
}
