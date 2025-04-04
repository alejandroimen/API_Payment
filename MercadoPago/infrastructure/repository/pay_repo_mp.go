package repository

import (
	"context"
	"fmt"

	"github.com/alejandroimen/API_Payment/MercadoPago/domain/entities"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

type PayRepoMP struct {
	cfg *config.Config
}

// Constructor del repositorio
func NewPayRepoMP(cfg *config.Config) *PayRepoMP {
	return &PayRepoMP{cfg: cfg}
}

func (r *PayRepoMP) ProccessPayment(data entities.DataPayment) (*payment.Response, error) {

	client := payment.NewClient(r.cfg)

	request := payment.Request{
		TransactionAmount: float64(data.TransactionAmount),
		Payer: &payment.PayerRequest{
			Email: data.Email,
		},
		Token:        data.Token,
		Installments: int(data.Installments),
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(resource)
	return resource, err
}
