package main

import (
	"net/http"
)

func main() {

	dbCreate(dbFile, dbCreateSQL)

	// http.HandleFunc("/duplicates/", Duplicates)
	http.Handle("/", http.FileServer(http.Dir("/app/html")))
	http.ListenAndServe(":8080", nil)
}
