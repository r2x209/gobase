package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	FullName string
	Email    string
	Password string
	Active   string
}

func createUser(user User) int {

	db, _ := sql.Open("sqlite3", databaseDSN)
	defer db.Close()

	row := db.QueryRow(
		"INSERT INTO Users (FullName, Email, Password, Active) "+
			"VALUES ($1, $2, $3, $4) RETURNING ID",
		user.FullName, user.Email, user.Password, user.Active,
	)

	var ID int64
	row.Scan(&ID)

	return int(ID)
}

func readUser(ID int) User {

	db, _ := sql.Open("sqlite3", databaseDSN)
	defer db.Close()

	row := db.QueryRow("SELECT * FROM Users WHERE ID=$1", ID)

	var user User
	err := row.Scan(&user.ID, &user.FullName, &user.Email, &user.Password, &user.Active)

	if err != nil {
		user.ID = 0
		user.FullName = ""
		user.Email = ""
		user.Password = ""
		user.Active = "Y"
	}

	return user
}

func updateUser(user User) {

}

func deleteUser(ID int) {

}

func listUser() []User {

	db, _ := sql.Open("sqlite3", databaseDSN)
	defer db.Close()

	rs, _ := db.Query("SELECT * FROM Users ORDER BY FullName")
	defer rs.Close()

	var users []User
	for rs.Next() {
		var r User
		rs.Scan(&r.ID, &r.FullName, &r.Email, &r.Password, &r.Active)
		users = append(users, r)
	}

	return users
}

// ---------------------------------------------------------------------------------------------------------------------

func createUserForm(w http.ResponseWriter, r *http.Request) {

	type TemplateData struct {
		AppTitle  string
		PageTitle string
		User      User
	}

	var user User
	user.FullName = ""
	user.Email = ""
	user.Password = ""
	user.Active = "Y"

	var tplData TemplateData
	tplData.AppTitle = appTitle
	tplData.PageTitle = "New User"
	tplData.User = user

	tpl, _ := template.ParseFS(tplDir, "templates/staff_user.html", "templates/staff_base.html")
	tpl.ExecuteTemplate(w, "staff_user.html", tplData)
}

func createUserHTML(w http.ResponseWriter, r *http.Request) {

	active := "N"
	if r.FormValue("Active") == "on" {
		active = "Y"
	}

	var user User
	user.FullName = r.FormValue("FullName")
	user.Email = r.FormValue("Email")
	user.Password = EncodeMD5(r.FormValue("Password"))
	user.Active = active

	fmt.Fprintln(w, createUser(user))
}

func deleteUserHTML(w http.ResponseWriter, r *http.Request) {

}

func readUserHTML(w http.ResponseWriter, r *http.Request) {

	type TemplateData struct {
		AppTitle  string
		PageTitle string
		User      User
	}

	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["ID"])

	var tplData TemplateData
	tplData.AppTitle = appTitle
	tplData.PageTitle = "Users"
	tplData.User = readUser(ID)

	tpl, _ := template.ParseFS(tplDir, "templates/staff_user.html", "templates/staff_base.html")
	tpl.ExecuteTemplate(w, "staff_user.html", tplData)
}

func listUserHTML(w http.ResponseWriter, r *http.Request) {

	type TemplateData struct {
		AppTitle  string
		PageTitle string
		Users     []User
	}

	var tplData TemplateData
	tplData.AppTitle = appTitle
	tplData.PageTitle = "Users"
	tplData.Users = listUser()

	tpl, _ := template.ParseFS(tplDir, "templates/staff_users.html", "templates/staff_base.html")
	tpl.ExecuteTemplate(w, "staff_users.html", tplData)
}
