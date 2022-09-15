package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/files"
	"io/ioutil"
	"sort"
)

type Service struct {
	Name string
	Hits int
}

func Services() {
	sampleFiles, err := ioutil.ReadDir("samples")
	if err != nil {
		fmt.Println(err)
		return
	}

	servicesMap := map[string]int{}

	for _, file := range sampleFiles {
		jsonString := files.ReadFile("samples/" + file.Name())
		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)
		for _, d := range logResponse.Data {
			servicesMap[d.Attributes.Service]++
		}
	}

	services := []Service{}

	for k, v := range servicesMap {
		s := Service{k, v}
		services = append(services, s)
	}

	sort.SliceStable(services, func(i, j int) bool {
		return services[i].Hits > services[j].Hits
	})

	for i, service := range services {
		fmt.Printf("%03d. %-60s %d\n", i+1, service.Name, service.Hits)
	}

}
