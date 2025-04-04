package routes

import (
	"github.com/alejandroimen/API_Payment.git/MercadoPago/infrastructure/controllers"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupRoutes(
	engine *gin.Engine,
	createSoapController *controllers.
) {

	// Grupo de rutas públicas ( cambiar despues a protegidas porfa)
	engine.POST("/pay", createSoapController.Handle)
}