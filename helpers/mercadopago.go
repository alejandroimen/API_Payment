package helpers

import (
	"fmt"
	"os"
	"github.com/mercadopago/sdk-go/pkg/config"
)


func NewMercadoPagoSDK() (*config.Config, error) {
	accessToken := os.Getenv("ACCESS_TOKEN")

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cfg, err
}