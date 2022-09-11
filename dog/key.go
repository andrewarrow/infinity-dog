package dog

import (
	"fmt"
	"infinity-dog/network"
)

func CheckKey() {
	jsonString := network.DoGet("/api/v1/validate")
	fmt.Println(jsonString)
}
