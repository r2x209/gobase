package main

import (
	"database/sql"
	"embed"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed templates
var tplDir embed.FS

//go:embed static
var staticDir embed.FS

func createDB() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	db.Exec(dbSchema)
	db.Exec(dbData)
}

func runServer() {
	http.Handle("/static/", http.FileServer(http.FS(staticDir)))

	http.HandleFunc("/api/services/", getAllServicesJSON)
	http.HandleFunc("/api/services/available/", getAvailableServicesJSON)

	http.HandleFunc("/", getAvailableServicesHTML)
	http.HandleFunc("/services/", getAvailableServicesHTML)

	http.ListenAndServe(":8080", nil)
}

func main() {
	createDB()
	runServer()
}
