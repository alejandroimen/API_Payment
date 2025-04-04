package application

import (
	"github.com/alejandroimen/API_Payment/MercadoPago/domain/entities"
	"github.com/alejandroimen/API_Payment/MercadoPago/domain/repository"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

type Pay struct {
	repo repository.PayRepoMP
}

func NewPay(repo repository.PayRepoMP) *Pay {
	return &Pay{repo: repo}
}

func (p *Pay) Run(transactionAmount float32, email string, token string, installments int16) (*payment.Response, error) {
	data := entities.DataPayment{
		TransactionAmount: transactionAmount,
		Email:             email,
		Token:             token,
		Installments:      installments,
	}

	Pay, err := p.repo.ProccessPayment(data)
	if err != nil {
		return nil, err
	}
	return Pay, nil
}
