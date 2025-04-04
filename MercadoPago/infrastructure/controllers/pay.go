package controllers

import (
	"log"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	"github.com/gin-gonic/gin"
)

type CreateSoapController struct {
	createSoap *application.CreateSoap
}

func NewCreateSoapController(createSoap *application.CreateSoap) *CreateSoapController {
	return &CreateSoapController{createSoap: createSoap}
}

func (c *CreateSoapController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para crear un jabón")

	var request struct {
		TransactionAmount float32 `json:"transaction_amount"`
		Email string `json:"email"`
		Token string `json:"token"`
		Installments int16 `json:"installments"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando el cuerpo de la solicitud: %v", err)
		ctx.JSON(400, gin.H{"error": "cuerpo de la solicitud inválido"})
		return
	}

	log.Printf("Creando jabón: Nombre=%s, Marca=%s, Tipo=%s, Precio=%f, Densidad=%f, Usuario id=%d",
		request.TransactionAmount, request.Email, request.Token, request.Installments)

	if err := c.createSoap.Run(request.TransactionAmount, request.Email, request.Token, request.Installments); err != nil {
		log.Printf("Error creando el jabón: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Paguito hecho unu")
	ctx.JSON(201, gin.H{"message": "pago realizado exitosamente"})
}