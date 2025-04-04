package repository

import (
	entities "github.com/alejandroimen/API-Payment.git/MercadoPago/domain/entities"
)

type MpRepository interface {
	ProccessPayment(data entities.DataPayment) error
} 