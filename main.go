package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	d, err := sql.Open("mysql", "root@tcp(db:3306)/")
	if err != nil {
		panic(err)
	}

	db = d

	_, err = db.Exec("CREATE DATABASE example")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE example")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE boards (
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		title varchar(5000) NOT NULL,
		description varchar(5000) NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("GET /boards", HandleList)
	http.HandleFunc("POST /boards", HandlePOST)
	http.HandleFunc("GET /boards/{id}", HandleGET)
	http.HandleFunc("PUT /boards/{id}", HandlePUT)
	http.HandleFunc("DELETE /boards/{id}", HandleDelete)

	http.ListenAndServe(":8080", nil)
}
