package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/negroni"
)

var db *sql.DB

type Info struct {
	ID         string `json:"id"`
	First_Name string `json:"firstname"`
	Instagram  string `json:"instagram"`
	File_Path  string `json:"filepath"`
}

type Infos []Info

func getAll(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM images")
	if err != nil {
		log.Fatal(err)
	}

	var images Infos

	for rows.Next() {
		var image Info
		rows.Scan(&image.ID, &image.First_Name, &image.Instagram, &image.File_Path)
		images = append(images, image)
	}
	jsonB, err := json.Marshal(images)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonB)

	//fmt.Fprintf(w, "%s", string(jsonB)) use fmt

}

func getId(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	stmt, err := db.Prepare(" SELECT * FROM images where ID == ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}

	var image Info

	for rows.Next() {
		rows.Scan(&image.ID, &image.First_Name, &image.Instagram, &image.File_Path)
	}

	jsonB, err := json.Marshal(image)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonB)

}

func One(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var image Info

	err := db.QueryRow(" SELECT * FROM images where ID == ?", vars["id"]).Scan(&image.ID, &image.First_Name, &image.Instagram, &image.File_Path)
	jsonB, err := json.Marshal(image)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonB)

}

func main() {

	database, err := sql.Open("sqlite3", "./nexai.db")
	if err != nil {
		log.Fatal(err)
	}

	db = database

	mux := mux.NewRouter()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome! Getting started!! \n"))
	}).Methods("GET")

	mux.HandleFunc("/all", getAll).Methods("GET")
	mux.HandleFunc("/image/{id:[0-9]+}", getId).Methods("GET")
	mux.HandleFunc("/img/{id:[0-9]+}", One).Methods("GET")

	// Includes some default middlewares
	n := negroni.Classic()
	n.UseHandler(mux)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
