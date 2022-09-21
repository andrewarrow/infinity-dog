package main

import (
	"infinity-dog/dog"
	"infinity-dog/screen"
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
		dog.CreateKey()
	} else if command == "billing" {
		dog.Billing()
	} else if command == "import" {
		dog.Import()
	} else if command == "sample" {
		hours := os.Args[2]
		dog.Sample(hours)
	} else if command == "exceptions" {
		service := util.GetArg(2)
		dog.Exceptions(service)
	} else if command == "messages" {
		service := util.GetArg(2)
		dog.Messages(service)
	} else if command == "screen" {
		screen.Setup()
	} else if command == "services" {
		sort := util.GetArg(2)
		level := util.GetArg(3)
		_ = level
		dog.ServicesFromSql(sort, "")
	} else if command == "help" {
		PrintHelp()
	}
}
