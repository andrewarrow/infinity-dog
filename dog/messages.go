package dog

import (
	"fmt"
	"sort"
	"strings"
)

type Message struct {
	Name string
	Hits int
}

func Messages(service string) {
	messagesMap := map[string]int{}
	ServicesFromSql("", service)

	for _, m := range servicesMessages {
		bothMessages := m
		if len(bothMessages) > 162 {
			bothMessages = bothMessages[:162]
		}
		tokens := strings.Split(bothMessages, "\n")
		messagesMap[tokens[0]]++
	}

	messagesList := []Message{}
	for k, v := range messagesMap {
		m := Message{}
		m.Name = k
		m.Hits = v
		messagesList = append(messagesList, m)
	}
	sort.SliceStable(messagesList, func(i, j int) bool {
		return messagesList[i].Hits > messagesList[j].Hits
	})

	for i, messages := range messagesList {
		fmt.Printf("%03d. [%03d] %-60s\n", i+1, messages.Hits, messages.Name)
	}

}
