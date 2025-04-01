package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/API-1-Multi/Order/application"
)

type CreateOrderController struct {
	createOrder *application.CreateOrder
}

func NewCreateOrderController(createOrder *application.CreateOrder) *CreateOrderController {
	return &CreateOrderController{createOrder: createOrder}
}

func (c *CreateOrderController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para crear una orden")

	var request struct {
		Cantidad            float64 `json:"cantidad"`
		Estado              int     `json:"estado"`
		Costo               float64 `json:"costo"`
		Codigo_Identificador string  `json:"codigo_identificador"`
		Tipo                bool    `json:"tipo"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cuerpo de la solicitud inválido"})
		return
	}

	log.Printf("Creando orden: Cantidad=%f, Estado=%d, Costo=%f, Codigo_Identificador=%s, Tipo=%t",
		request.Cantidad, request.Estado, request.Costo, request.Codigo_Identificador, request.Tipo)

	if err := c.createOrder.Run(request.Cantidad, request.Estado, request.Costo, request.Codigo_Identificador, request.Tipo); err != nil {
		log.Printf("Error creando la orden: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando la orden"})
		return
	}

	log.Println("Orden creada exitosamente")
	ctx.JSON(http.StatusCreated, gin.H{"message": "Orden creada exitosamente"})
}