package controllers

import (
    "log"
    "net/http"

    "github.com/alejandroimen/API_Payment/MercadoPago/application"
    "github.com/gin-gonic/gin"
)

type PayController struct {
    payUseCase *application.Pay
}

func NewPayController(payUseCase *application.Pay) *PayController {
    return &PayController{payUseCase: payUseCase}
}

func (c *PayController) Handle(ctx *gin.Context) {
    // Modificar la estructura de datos esperada para incluir `payer`
    var req struct {
        Payer struct {
            Email string `json:"email"`
        } `json:"payer"`
        Token string `json:"token"`
    }

    // Intentar enlazar el cuerpo JSON con la estructura
    if err := ctx.ShouldBindJSON(&req); err != nil {
        log.Println("Error decodificando el cuerpo:", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cuerpo de la solicitud inválido"})
        return
    }

    // Extraer el email del subcampo `payer`
    email := req.Payer.Email
    token := req.Token

    log.Printf("Solicitud recibida: Email=%s, Token=%s", email, token)

    // Pasar los datos al caso de uso
    resp, err := c.payUseCase.Run(email, token)
    if err != nil {
        log.Println("Error ejecutando el caso de uso:", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    log.Println("Respuesta del caso de uso:", resp)

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Pago realizado con éxito",
        "cuerpo":  resp,
    })
}
