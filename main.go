package main

import (
	"database/sql"
	"embed"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed templates
var tplDir embed.FS

//go:embed static
var staticDir embed.FS

func createDB() {
	db, err := sql.Open("sqlite3", databaseDSN)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	db.Exec(databaseSchema)

	var users int
	row := db.QueryRow("SELECT COUNT(*) FROM Users")
	err = row.Scan(&users)
	if err != nil {
		log.Fatal(err.Error())
	}

	if users == 0 {
		db.Exec(databaseData)
	}
}

// func runServer() {
// 	http

// 	mux.Handle("/static/", http.FileServer(http.FS(staticDir)))

// 	// http.HandleFunc("/api/services/", getAllServicesJSON)
// 	// http.HandleFunc("/api/services/available/", getAvailableServicesJSON)

// 	// http.HandleFunc("/", getAvailableServicesHTML)
// 	// http.HandleFunc("/services/", getAvailableServicesHTML)

// 	http.HandleFunc("/staff/users/", getUsersHTML)
// 	http.HandleFunc("/staff/user/", getUserHTML)

// 	http.ListenAndServe(":8080", nil)
// }

func serve() {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/staff/user/list", listUserHTML).Methods("GET")
	r.HandleFunc("/staff/user", createUserForm).Methods("GET")
	r.HandleFunc("/staff/user", createUserHTML).Methods("POST")
	r.HandleFunc("/staff/user/{ID}", readUserHTML).Methods("GET")
	r.HandleFunc("/staff/user/{ID}", deleteUserHTML).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {

	createDB()

	serve()
}
