package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func checkoutPage(w http.ResponseWriter, r *http.Request) {
	// var t payment.Token

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	fmt.Printf("body: %+v", body)

	// err := json.NewDecoder(r.Body).Decode(&t)
	// if err != nil {
	// 	panic(err)
	// }

	// payment.Checkout(t)
}

func main() {
	http.HandleFunc("/checkout", checkoutPage)
	http.HandleFunc("/", indexPage)

	log.Print("running server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
