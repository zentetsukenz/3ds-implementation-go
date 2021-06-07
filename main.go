package main

import (
	"html/template"
	"log"
	"net/http"
)

func payPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", payPage)

	log.Print("running server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
