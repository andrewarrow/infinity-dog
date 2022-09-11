package main

import (
	"infinity-dog/dog"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "logs" {
		query := os.Args[2]
		dog.Logs(query)
	} else if command == "key" {
		dog.CheckKey()
	} else if command == "help" {
		PrintHelp()
	}
}
