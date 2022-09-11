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
		dog.Logs()
	} else if command == "help" {
		PrintHelp()
	}
}
