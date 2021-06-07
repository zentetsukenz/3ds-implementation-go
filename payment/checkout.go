package payment

import (
	"fmt"
	"log"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

const (
	OmisePublicKey = "pkey_test_"
	OmiseSecretKey = "skey_test_"
)

type Token struct {
	OmiseToken  string
	OmiseSource string
}

func Checkout(token Token) (*omise.Charge, error) {
	log.Printf("got token %+v", token)

	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Fatal(e)
		return nil, e
	}

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:      100000,
		Card:        token.OmiseToken,
		Currency:    "thb",
		Description: "Test payment from som-m/3ds-implementation-go",
		ReturnURI:   fmt.Sprintf("http://localhost:8080/complete?order_id=%s", token.OmiseToken),
		Metadata: map[string]interface{}{
			"order_id": token.OmiseToken,
		},
	}

	if e := client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
		return nil, e
	}

	log.Printf("charge: %+v\n", charge)

	return charge, nil
}
