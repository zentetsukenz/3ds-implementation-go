package main

import (
	"go/payment"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/omise/omise-go"
)

var decoder = schema.NewDecoder()

func indexPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func checkoutPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)

		return
	}

	var (
		t      payment.Token
		charge *omise.Charge
	)

	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		log.Printf("Error decoding body: %v", err)
		http.Error(w, "can't decode body", http.StatusInternalServerError)

		return
	}

	charge, err = payment.Checkout(t)
	if err != nil {
		log.Printf("Error making a charge: %v", err)
		http.Error(w, "cannot charge", http.StatusBadGateway)

		return
	}

	http.Redirect(w, r, charge.AuthorizeURI, http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/checkout", checkoutPage)
	http.HandleFunc("/", indexPage)

	log.Print("running server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
