package routes

import (
	"github.com/alejandroimen/API_Payment/MercadoPago/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	engine *gin.Engine,
	payController *controllers.PayController,
) {

	// Grupo de rutas públicas ( cambiar despues a protegidas porfa)
	engine.POST("/pay", payController.Handle)
}