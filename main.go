package main

import (
	"embed"
	"net/http"
)

//go:embed templates
var tplDir embed.FS

//go:embed static
var staticDir embed.FS

func main() {

	http.Handle("/static/", http.FileServer(http.FS(staticDir)))

	createDB()

	http.HandleFunc("/api/services/", getAllServicesJSON)
	http.HandleFunc("/api/services/available/", getAvailableServicesJSON)

	http.HandleFunc("/", getAvailableServicesHTML)
	http.HandleFunc("/services/", getAvailableServicesHTML)

	http.ListenAndServe(":8080", nil)
}
