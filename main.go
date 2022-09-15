package main

import (
	"infinity-dog/dog"
	"infinity-dog/util"
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
	os.Mkdir("samples", 0755)

	if command == "logs" {
		query := os.Args[2]
		dog.Logs(24, query)
	} else if command == "key" {
		dog.CheckKey()
	} else if command == "sample" {
		hours := os.Args[2]
		dog.Sample(hours)
	} else if command == "services" {
		sort := util.GetArg(2)
		level := util.GetArg(3)
		dog.Services(sort, level)
	} else if command == "help" {
		PrintHelp()
	}
}
