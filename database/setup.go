package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenTheDB() *sql.DB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func CreateSchema() {
	db := OpenTheDB()
	defer db.Close()

	sqlStmt := `
CREATE TABLE services (name text, msg text, message text, exception text, logged_at datetime);
CREATE INDEX IF NOT EXISTS index1 ON services (name);
CREATE INDEX IF NOT EXISTS index2 ON services (logged_at);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q\n", err)
		return
	}
}
