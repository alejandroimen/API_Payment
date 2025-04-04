package application

import (
	"github.com/alejandroimen/API-Payment.git/MercadoPago/domain/entities"
	"github.com/alejandroimen/API-Payment.git/MercadoPago/domain/repository"
	"github.com/alejandroimen/API_Payment.git/MercadoPago/domain/entities"
)

type pay struct {
	repo repository.MpRepository
}

func NewPay(repo repository.MpRepository) *pay{
	return &pay{repo: repo}
}

func(p *pay) Run(transactionAmount float32, email string, token string, installments int16) error {
	data := entities.DataPayment{
		TransactionAmount: transactionAmount, 
		Email: email,
		Token: token,
		Installments: installments,
	}

	pay, err := p.repo.ProccessPayment(data)
	if err != nil {
		return nil, err
	}
	return pay, nil
}