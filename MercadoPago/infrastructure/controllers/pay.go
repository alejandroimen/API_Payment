package controllers

import (
	"log"

	"github.com/alejandroimen/API_Payment/MercadoPago/application"
	"github.com/gin-gonic/gin"
)

type PayController struct {
	pay *application.Pay
}

func NewPayController(pay *application.Pay) *PayController {
	return &PayController{pay: pay}
}

func (c *PayController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para crear un jabón")

	var request struct {
		TransactionAmount float32 `json:"transaction_amount"`
		Email             string  `json:"email"`
		Token             string  `json:"token"`
		Installments      int16   `json:"installments"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando el cuerpo de la solicitud: %v", err)
		ctx.JSON(400, gin.H{"error": "cuerpo de la solicitud inválido"})
		return
	}

	log.Printf("Creando pago: Monto=%f, Email=%s, Token=%s, Installment=%d",
		request.TransactionAmount, request.Email, request.Token, request.Installments)

	pay, err := c.pay.Run(request.TransactionAmount, request.Email, request.Token, request.Installments); 
	if err != nil {
		log.Printf("Error creando el jabón: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Paguito hecho unu", pay)
	ctx.JSON(201, gin.H{"message": "pago realizado exitosamente"})
}
