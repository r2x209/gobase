package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func dbCreate(filename string, statements []string) {

	if _, err := os.Stat(filename); err != nil {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
	}

	for _, value := range statements {
		dbExecute(value)
	}
}

func dbExecute(statement string) {

	db, _ := sql.Open("sqlite3", dbFile)
	defer db.Close()

	query, err := db.Prepare(statement)
	if err != nil {
		log.Fatal(err.Error())
	}

	query.Exec()
}
