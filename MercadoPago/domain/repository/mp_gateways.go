package repository

import "github.com/alejandroimen/API_Payment/MercadoPago/domain/entities"

type PayRepo interface {
    CreatePayment(payment entities.Payment) (interface{}, error)
}
