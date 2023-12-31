package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Duration    int     `json:"duration"`
	Price       float32 `json:"price"`
	Available   string  `json:"available"`
}

type TemplateData struct {
	AppTitle  string
	PageTitle string
	Services  []Service
}

func getServices(query string) []Service {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	defer db.Close()

	recordset, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer recordset.Close()

	var services []Service

	for recordset.Next() {
		var s Service
		recordset.Scan(&s.ID, &s.Title, &s.Description, &s.Duration, &s.Price, &s.Available)
		services = append(services, s)
	}

	return services
}

func getAllServicesJSON(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM Services ORDER BY Title ASC"
	services := getServices(query)
	json.NewEncoder(w).Encode(services)
}

func getAvailableServicesJSON(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM Services WHERE Active='Y' ORDER BY Title ASC"
	services := getServices(query)
	json.NewEncoder(w).Encode(services)
}

func getAvailableServicesHTML(w http.ResponseWriter, r *http.Request) {
	var tplData TemplateData
	tplData.AppTitle = appTitle
	tplData.PageTitle = "Services"

	query := "SELECT * FROM Services WHERE Active='Y' ORDER BY title ASC"
	tplData.Services = getServices(query)

	tmpl, _ := template.ParseFS(tplDir, "templates/services.html", "templates/base.html")
	tmpl.ExecuteTemplate(w, "services.html", tplData)
}
