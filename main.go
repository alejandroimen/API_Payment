package main

import (
	"log"

	mp_infra "github.com/alejandroimen/API_Payment/MercadoPago/infrastructure"

	"github.com/alejandroimen/API_Payment/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := helpers.NewMercadoPagoSDK()
	if err != nil {
		log.Fatalf("Error inicializando la conexión a MySQL: %v", err)
	}

	engine := gin.Default()
	engine.Use(helpers.SetupCORS())

	mp_infra.InitMpDependencies(engine, cfg)

	engine.Run(":8000")

}
