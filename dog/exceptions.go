package dog

import (
	"fmt"
	"sort"
	"strings"
)

type Exception struct {
	Name string
	Hits int
}

func Exceptions(service string) {
	exceptionMap := map[string]int{}
	ServicesFromSql("", service)

	for _, e := range servicesExceptions {
		tokens := strings.Split(e, "\n")
		exceptionMap[tokens[0]]++
	}

	exceptionList := []Exception{}
	for k, v := range exceptionMap {
		e := Exception{}
		e.Name = k
		e.Hits = v
		exceptionList = append(exceptionList, e)
	}
	sort.SliceStable(exceptionList, func(i, j int) bool {
		return exceptionList[i].Hits > exceptionList[j].Hits
	})

	for i, exception := range exceptionList {
		fmt.Printf("%03d. [%03d] %-60s\n", i+1, exception.Hits, exception.Name)
	}

}
