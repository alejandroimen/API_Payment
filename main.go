package main

import (
    "log"
    "os"

    mp_infra "github.com/alejandroimen/API_Payment/MercadoPago/infrastructure"
    "github.com/alejandroimen/API_Payment/helpers"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Cargar las variables de entorno desde el archivo .env
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error cargando archivo .env: %v", err)
    }

    // Verificar que las variables esenciales estén cargadas
    if os.Getenv("ACCESS_TOKEN") == "" {
        log.Fatal("ACCESS_TOKEN no está definido en las variables de entorno")
    }

    // Inicializar SDK de MercadoPago
    cfg, err := helpers.NewMercadoPagoSDK()
    if err != nil {
        log.Fatalf("Error inicializando el SDK de MercadoPago: %v", err)
    }

    // Crear instancia del router
    engine := gin.Default()

    // Configurar middlewares (e.g., CORS)
    engine.Use(helpers.SetupCORS())

    // Inicializar las dependencias, pasando el SDK de MercadoPago
    mp_infra.InitMpDependencies(engine, cfg)

    // Iniciar el servidor en el puerto 8000
    if err := engine.Run(":8000"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}
