package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/files"
	"io/ioutil"
	"sort"
	"strings"
)

type Exception struct {
	Name string
	Hits int
}

func Exceptions(service string) {
	sampleFiles, err := ioutil.ReadDir("samples")
	if err != nil {
		fmt.Println(err)
		return
	}

	exceptionMap := map[string]int{}

	for _, file := range sampleFiles {
		jsonString := files.ReadFile("samples/" + file.Name())
		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)
		for _, d := range logResponse.Data {
			if d.Attributes.Service != service {
				continue
			}
			if len(d.Attributes.SubAttributes.Exception) == 0 {
				continue
			}
			e := d.Attributes.SubAttributes.Exception
			tokens := strings.Split(e, "\n")
			exceptionMap[tokens[0]]++
		}
	}

	exceptionList := []Exception{}
	for k, v := range exceptionMap {
		s := Exception{}
		s.Name = k
		s.Hits = v
		exceptionList = append(exceptionList, s)
	}
	sort.SliceStable(exceptionList, func(i, j int) bool {
		return exceptionList[i].Hits > exceptionList[j].Hits
	})

	for i, exception := range exceptionList {
		fmt.Printf("%03d. [%03d] %-60s\n", i+1, exception.Hits, exception.Name)
	}

}
