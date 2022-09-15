package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/files"
	"io/ioutil"
)

func Services() {
	sampleFiles, err := ioutil.ReadDir("samples")
	if err != nil {
		fmt.Println(err)
		return
	}

	services := map[string]int{}

	for _, file := range sampleFiles {
		jsonString := files.ReadFile("samples/" + file.Name())
		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)
		for _, d := range logResponse.Data {
			services[d.Attributes.Service]++
		}
	}

	for k, v := range services {
		fmt.Println(k, v)
	}
}
