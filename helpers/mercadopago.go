package helpers

import (
	"os",
	"github.com/mercadopago/sdk-go/pkg/config"
)


func newMercadoPagoSDK() {
	accessToken := os.Getenv("ACCESS_TOKEN")

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	return cfg, err
}