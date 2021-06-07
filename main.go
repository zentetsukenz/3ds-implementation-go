package main

import (
	"log"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

const (
	OmisePublicKey = "pkey_test_"
	OmiseSecretKey = "skey_test_"
)

func main() {
	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Fatal(e)
	}

	card, createToken := &omise.Card{}, &operations.CreateToken{
		Name:            "Test Test",
		Number:          "4242424242424242",
		ExpirationMonth: 10,
		ExpirationYear:  2022,
		SecurityCode:    "123",
	}

	if e := client.Do(card, createToken); e != nil {
		log.Fatalln(e)
	}

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   100000,
		Currency: "thb",
		Card:     card.ID,
	}
	if e := client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}

	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
}