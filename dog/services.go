package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/database"
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

func ServicesFromSql(sortString string) {
	s := `select distinct(name) as name, count(name) as c from services group by name order by c desc`
	db := database.OpenTheDB()
	defer db.Close()

	rows, err := db.Query(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		var t1 string
		var t2 string
		err = rows.Scan(&t1, &t2)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%03d. %-60s %s\n", i+1, t1, t2)
		i++
	}
}

func Services(sortString, level string) {
	sampleFiles, err := ioutil.ReadDir("samples")
	if err != nil {
		fmt.Println(err)
		return
	}

	servicesHitMap := map[string]int{}
	servicesDataMap := map[string]int{}
	servicesExceptionMap := map[string]int{}

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
			if len(d.Attributes.SubAttributes.Exception) > 0 {
				servicesExceptionMap[d.Attributes.Service]++
			}
		}
	}

	servicesList := []Service{}

	for k, v := range servicesHitMap {
		s := Service{}
		s.Name = k
		s.Hits = v
		s.Data = servicesDataMap[k]
		s.Exceptions = servicesExceptionMap[k]
		servicesList = append(servicesList, s)
	}

	theSort := "hits"
	if sortString != "" {
		theSort = sortString
	}

	if theSort == "hits" {
		sort.SliceStable(servicesList, func(i, j int) bool {
			return servicesList[i].Hits > servicesList[j].Hits
		})
	} else if theSort == "data" {
		sort.SliceStable(servicesList, func(i, j int) bool {
			return servicesList[i].Data > servicesList[j].Data
		})
	} else if theSort == "exceptions" {
		sort.SliceStable(servicesList, func(i, j int) bool {
			return servicesList[i].Exceptions > servicesList[j].Exceptions
		})
	}

	if theSort == "hits" {
		for i, service := range servicesList {
			fmt.Printf("%03d. %-60s %d\n", i+1, service.Name, service.Hits)
		}
	} else if theSort == "data" {
		for i, service := range servicesList {
			fmt.Printf("%03d. %-60s %d\n", i+1, service.Name, service.Data)
		}
	} else if theSort == "exceptions" {
		for i, service := range servicesList {
			fmt.Printf("%03d. %-60s %d\n", i+1, service.Name, service.Exceptions)
		}
	}

}
