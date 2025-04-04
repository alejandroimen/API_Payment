package infrastructure

import (
	"github.com/mercadopago/sdk-go/pkg/config"

	app_mp "github.com/alejandroimen/API_Payment/MercadoPago/application"
	control_mp "github.com/alejandroimen/API_Payment/MercadoPago/infrastructure/controllers"
	repo_mp "github.com/alejandroimen/API_Payment/MercadoPago/infrastructure/repository"
	routes_mp "github.com/alejandroimen/API_Payment/MercadoPago/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func InitMpDependencies(engine *gin.Engine, cfg *config.Config) {

	mpRepo := repo_mp.NewPayRepoMP(cfg)

	creatempUseCase := app_mp.NewPay(mpRepo)

	creatempController := control_mp.NewPayController(creatempUseCase)

	routes_mp.SetupRoutes(
		engine,
		creatempController,
	)

}
