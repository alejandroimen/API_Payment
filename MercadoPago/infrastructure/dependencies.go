package infrastructure

import (
    "github.com/alejandroimen/API_Payment/MercadoPago/application"
    "github.com/alejandroimen/API_Payment/MercadoPago/infrastructure/controllers"
    "github.com/alejandroimen/API_Payment/MercadoPago/infrastructure/repository"
    "github.com/alejandroimen/API_Payment/MercadoPago/infrastructure/routes"
    "github.com/gin-gonic/gin"
    "github.com/mercadopago/sdk-go/pkg/config"
)

func InitMpDependencies(engine *gin.Engine, cfg *config.Config) {
    repo := repository.NewPayRepoMP(cfg) // Ahora se pasa cfg al repositorio
    useCase := application.NewPay(repo)
    controller := controllers.NewPayController(useCase)

    routes.SetupRoutes(engine, controller)
}
