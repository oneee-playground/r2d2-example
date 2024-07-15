package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostData struct {
	Title       string `json:"titles"`
	Description string `json:"description"`
}

func HandlePOST(w http.ResponseWriter, r *http.Request) {
	var data PostData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	query := fmt.Sprintf(
		"INSERT INTO boards(title, description) VALUES ('%s', '%s')", data.Title, data.Description,
	)

	if _, err := db.Exec(query); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
}

func HandlePUT(w http.ResponseWriter, r *http.Request) {
	var data PostData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	id := r.PathValue("id")

	query := fmt.Sprintf(
		"UPDATE boards SET title = %s, description = %s WHERE id = %s", data.Title, data.Description, id,
	)

	if _, err := db.Exec(query); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func HandleGET(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	query := fmt.Sprintf(
		"SELECT * FROM boards WHERE id = %s", id,
	)

	var theID string
	var title string
	var description string

	err := db.QueryRow(query).Scan(&theID, &title, &description)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	b, _ := json.Marshal(map[string]string{
		"id":          theID,
		"title":       title,
		"description": description,
	})

	w.Write(b)

	w.WriteHeader(200)
}

func HandleList(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, title FROM boards ofder by id asc LIMIT 20"

	rows, err := db.Query(query)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	type Result struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}

	list := make([]Result, 0)
	for rows.Next() {
		var dst Result
		if err := rows.Scan(&dst.ID, &dst.Title); err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}

		list = append(list, dst)
	}

	b, _ := json.Marshal(list)

	w.Write(b)

	w.WriteHeader(200)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	query := fmt.Sprintf(
		"DELETE FROM boards WHERE id = %s", id,
	)

	if _, err := db.Exec(query); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
