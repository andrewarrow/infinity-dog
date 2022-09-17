package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/database"
	"infinity-dog/files"
	"io/ioutil"
	"sort"
	"strconv"
)

var servicesHitMap = map[string]int{}
var servicesDataMap = map[string]int{}
var servicesExceptionMap = map[string]int{}
var servicesExceptions = []string{}
var servicesMessages = []string{}

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

func ServicesHitsFromSql() []Service {
	items := []Service{}
	s := `select distinct(name) as name, count(name) as c from services group by name order by c desc`
	db := database.OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	for rows.Next() {
		var name string
		var hits string
		rows.Scan(&name, &hits)
		service := Service{}
		service.Name = name
		service.Hits, _ = strconv.Atoi(hits)
		items = append(items, service)
	}

	return items
}

func ServicesFromSql(sortString, service string) {
	//s := `select name, msg, message, exception from services order by logged_at`
	s := `select name, msg, message, exception from services`
	if service != "" {
		s = fmt.Sprintf(`select name, msg, message, exception from services where name='%s'`, service)
	}
	db := database.OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	for rows.Next() {
		var name string
		var msg string
		var message string
		var exception string
		rows.Scan(&name, &msg, &message, &exception)
		servicesHitMap[name]++
		dataLength := len(message) + len(msg) + len(exception)
		servicesDataMap[name] += dataLength
		if len(exception) > 0 {
			servicesExceptionMap[name]++
			servicesExceptions = append(servicesExceptions, exception)
		}
		if len(msg)+len(message) > 0 {
			servicesMessages = append(servicesMessages, message+msg)
		}
	}

	handleSortAndDisplay(sortString)
}

func handleSortAndDisplay(sortString string) {
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

func ServicesFromJson(sortString, level string) {
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

	handleSortAndDisplay(sortString)

}
