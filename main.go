package main

import (
	"net/http"

	_ "github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":3000", nil)
}
