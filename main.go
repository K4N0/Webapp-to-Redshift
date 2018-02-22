package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Hello, Go Web Development")
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":3000", nil)

}
