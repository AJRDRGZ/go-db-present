package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("driverName", "dataSourceName") // HL
	if err != nil {
		panic(err)
	}
	defer db.Close() // HL

	if err := db.Ping(); err != nil { // HL
		panic(err)
	}
}
