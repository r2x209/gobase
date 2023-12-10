package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbSchema string = `
	CREATE TABLE IF NOT EXISTS services 
	(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		title TEXT, 
		description TEXT,  
		duration INTEGER,  
		price NUMERIC, 
		available TEXT
	);

	CREATE TABLE IF NOT EXISTS orders 
	(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		date_time TEXT
	);

	CREATE TABLE IF NOT EXISTS customers 
	(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		name TEXT, 
		email TEXT, 
		address TEXT, 
		created TEXT
	);
`

var dbData string = `
	INSERT INTO services (title, description, duration, price, available)
		VALUES("Tire Rotation", "Take off current tires and put on alternative tires.", 30, 59.99, "y");
	INSERT INTO services (title, description, duration, price, available)
		VALUES("Tire Rotation with replacement", "Take off current tires and put on alternative tires.", 60, 99.99, "y");
`

func createDB() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	db.Exec(dbSchema)
	db.Exec(dbData)
}
