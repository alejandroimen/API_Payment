package repository

import (
	entities "github.com/alejandroimen/API_Payment/MercadoPago/domain/entities"
	"github.com/mercadopago/sdk-go/pkg/payment"

)

type PayRepoMP interface {
	ProccessPayment(data entities.DataPayment) (*payment.Response, error)
}
