package main

var dbFile string = "db.sqlite"

var dbCreateSQL = []string{
	"CREATE TABLE IF NOT EXISTS services (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, title TEXT, description TEXT,  duration INTEGER,  price NUMERIC, available TEXT)",
	"CREATE TABLE IF NOT EXISTS orders (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, date_time TEXT)",
	"CREATE TABLE IF NOT EXISTS customers (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT, address TEXT, created TEXT)",
}
