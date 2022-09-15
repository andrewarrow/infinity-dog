package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/files"
	"io/ioutil"
	"sort"
)

type Service struct {
	Name       string
	Hits       int
	Exceptions int
	Data       int
}

func NewService() *Service {
	s := Service{}
	return &s
}

func Services(sortString, level string) {
	sampleFiles, err := ioutil.ReadDir("samples")
	if err != nil {
		fmt.Println(err)
		return
	}

	servicesHitMap := map[string]int{}
	servicesDataMap := map[string]int{}

	for _, file := range sampleFiles {
		jsonString := files.ReadFile("samples/" + file.Name())
		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)
		for _, d := range logResponse.Data {
			servicesHitMap[d.Attributes.Service]++
			dataLength := len(d.Attributes.Message) +
				len(d.Attributes.SubAttributes.Msg) +
				len(d.Attributes.SubAttributes.Exception)
			servicesDataMap[d.Attributes.Service] += dataLength
		}
	}

	servicesList := []Service{}

	for k, v := range servicesHitMap {
		s := Service{}
		s.Name = k
		s.Hits = v
		s.Data = servicesDataMap[k]
		servicesList = append(servicesList, s)
	}

	sort.SliceStable(servicesList, func(i, j int) bool {
		//return services[i].Hits > services[j].Hits
		return servicesList[i].Data > servicesList[j].Data
	})

	for i, service := range servicesList {
		//fmt.Printf("%03d. %-60s %d\n", i+1, service.Name, service.Hits)
		fmt.Printf("%03d. %-60s %d\n", i+1, service.Name, service.Data)
	}

}
