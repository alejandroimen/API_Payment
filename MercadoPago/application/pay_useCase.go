package application

import (
    "log"
    "github.com/alejandroimen/API_Payment/MercadoPago/domain/entities"
    "github.com/alejandroimen/API_Payment/MercadoPago/domain/repository"
)

type Pay struct {
    repo repository.PayRepo
}

func NewPay(repo repository.PayRepo) *Pay {
    return &Pay{repo: repo}
}

func (p *Pay) Run(email string, token string) (interface{}, error) {
    data := entities.Payment{
        PayerEmail: email,
        Token:      token,
        PlanID:     "Plan mensual",
        Amount:     100.0,
        Currency:   "MXN",
    }
    log.Printf("Solicitud recibida: Email=%s, Token=%s", data.PayerEmail, data.Token)

    result, err := p.repo.CreatePayment(data)
    if err != nil {
        return nil, err
    }
    return result, nil
}
