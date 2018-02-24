package main

import (
	"database/sql"
	"html/template"
	"net/http"

	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

//Page blahblahblah
type Page struct {
	Name     string
	DBStatus bool
}

//Page blahblahblah1
type SearchResult struct {
	Title  string
	Author string
	Year   string
	ID     string
}

func main() {
	templates := template.Must(template.ParseFiles("index.html"))

	db, _ := sql.Open("sqlite3", "test.db")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "Gopher"}
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}

		p.DBStatus = db.Ping() == nil

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		results := []SearchResult{
			SearchResult{"Moby-Dick", "Herman Melville", "1851", "222222"},
			SearchResult{"The Adventures of Huckleberry Finn", "Mark Twain", "1884", "444444"},
			SearchResult{"The Catcher in the Rye", "JD Salinger", "1951", "333333"},
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":3000", nil)
}
