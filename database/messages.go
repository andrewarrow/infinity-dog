package database

import (
	"fmt"
)

type Message struct {
	Both     string
	LoggedAt string
}

func MessagesFromService(service string) []Message {
	items := []Message{}
	s := fmt.Sprintf(`select msg, message, logged_at from services where name='%s' order by logged_at desc limit 60`, service)

	db := OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	for rows.Next() {
		var msg string
		var messageString string
		var loggedAt string
		rows.Scan(&msg, &messageString, &loggedAt)
		message := Message{}
		message.Both = msg + messageString
		message.LoggedAt = loggedAt
		items = append(items, message)
	}

	return items
}
