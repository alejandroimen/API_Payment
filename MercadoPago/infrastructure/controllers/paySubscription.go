package controllers

import (
	"context"
	"fmt"

	"github.com/alejandroimen/API_Payment/helpers"

	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {

	cfg, err := helpers.newMercadoPagoSDK()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := payment.NewClient(cfg)

	request := payment.Request{
		TransactionAmount: 105.1,
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
		Token: "{{CARD_TOKEN}}",
		Installments: 1,
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}