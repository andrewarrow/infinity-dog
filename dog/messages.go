package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/files"
	"io/ioutil"
	"sort"
	"strings"
)

type Message struct {
	Name string
	Hits int
}

func Messages(service string) {
	sampleFiles, err := ioutil.ReadDir("samples")
	if err != nil {
		fmt.Println(err)
		return
	}

	messagesMap := map[string]int{}

	for _, file := range sampleFiles {
		jsonString := files.ReadFile("samples/" + file.Name())
		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)
		for _, d := range logResponse.Data {
			if d.Attributes.Service != service {
				continue
			}
			bothMessages := d.Attributes.Message + d.Attributes.SubAttributes.Msg
			if len(bothMessages) == 0 {
				continue
			}
			if len(bothMessages) > 162 {
				bothMessages = bothMessages[:162]
			}
			tokens := strings.Split(bothMessages, "\n")
			messagesMap[tokens[0]]++
		}
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
