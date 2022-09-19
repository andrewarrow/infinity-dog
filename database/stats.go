package database

import "fmt"

func TotalRows() int64 {
	s := fmt.Sprintf(`select count(1) from services`)

	db := OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	rows.Next()

	var total int64
	rows.Scan(&total)

	return total
}
