package main 

import (
	"github.com/alejandroimen/API_Payment/helpers"	
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := helpers.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error inicializando la conexión a MySQL: %v", err)
	}
	defer db.Close()

	engine := gin.Default()
	engine.Use(helpers.SetupCORS())
}