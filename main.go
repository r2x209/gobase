package main

import (
	"net/http"
)

func main() {
	createDB()

	http.HandleFunc("/api/services/", getAllServices)
	http.HandleFunc("/api/services/available/", getAvailableServices)

	http.Handle("/", http.FileServer(http.Dir("/app/html")))

	http.ListenAndServe(":8080", nil)
}
