package dog

import (
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

	for _, file := range sampleFiles {
		jsonString := files.ReadFile("samples/" + file.Name())
		fmt.Println(len(jsonString))

	}
}
