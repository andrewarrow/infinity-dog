package dog

import (
	"fmt"
	"infinity-dog/network"
)

func CheckKey() {
	jsonString := network.DoGet("/api/v1/validate")
	fmt.Println(jsonString)
}

func CreateKey() {
	input := `{
  "data": {
    "type": "application_keys",
    "attributes": {
      "name": "aa_usage_read5",
			"scopes": ["usage_read"]
    }
  }
}`
	jsonString := network.DoPost("/api/v2/current_user/application_keys", []byte(input))
	fmt.Println(jsonString)
}
