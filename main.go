package main

import (
	"net/http"
)

func main() {
	createDB()

	http.HandleFunc("/api/services/", getAllServicesJSON)
	http.HandleFunc("/api/services/available/", getAvailableServicesJSON)

	http.HandleFunc("/services/", getAvailableServicesHTML)

	http.Handle("/", http.FileServer(http.Dir("html")))

	http.ListenAndServe(":8080", nil)
}
