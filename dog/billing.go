package dog

import (
	"fmt"
	"infinity-dog/network"
)

func Billing() {
	// auth needs `usage_read` scope
	jsonString := network.DoGet("/api/v1/usage/billable-summary?month=2022-08")
	fmt.Println(jsonString)

}
