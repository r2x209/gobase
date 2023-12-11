package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createDB() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	db.Exec(dbSchema)
	db.Exec(dbData)
}
