package routes

import (
    "github.com/alejandroimen/API_Payment/MercadoPago/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine, controller *controllers.PayController) {
    engine.POST("/pay", controller.Handle)
}
