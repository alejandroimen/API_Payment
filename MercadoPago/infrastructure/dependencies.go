package infrastructure

import (
	"database/sql"
	"github.com/mercadopago/sdk-go/pkg/config"

	"github.com/gin-gonic/gin"
	app_mp "github.com/alejandroimen/API-Payment.git/MercadoPago/application"
	control_mp "github.com/alejandroimen/API-Payment.git/MercadoPago/infrastructure/controllers"
	repo_mp "github.com/alejandroimen/API-Payment.git/MercadoPago/infrastructure/repository"
	routes_mp "github.com/alejandroimen/API-Payment.git/MercadoPago/infrastructure/routes"
)

func InitMpDependencies(engine *gin.Engine, cfg *config.Config) {

	mpRepo := repo_mp.NewPayRepoMP(cfg)

	creatempUseCase := app_mp.New(mpRepo)

	creatempController := control_mp.NewCreatempController(creatempUseCase)

	routes_mp.SetupRoutes(
		engine,
		creatempController,
	)

}