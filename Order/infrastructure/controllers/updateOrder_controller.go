package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/API-1-Multi/Order/application"
)

type UpdateOrderController struct {
	updateOrder *application.UpdateOrder
}

func NewUpdateOrderController(updateOrder *application.UpdateOrder) *UpdateOrderController {
	return &UpdateOrderController{updateOrder: updateOrder}
}

func (u *UpdateOrderController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para actualizar una orden")

	var request struct {
		Id_Jabon            int     `json:"id_jabon"`
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

	log.Printf("Actualizando orden: Id_Jabon=%d, Cantidad=%f, Estado=%d, Costo=%f, Codigo_Identificador=%s, Tipo=%t",
		request.Id_Jabon, request.Cantidad, request.Estado, request.Costo, request.Codigo_Identificador, request.Tipo)

	if err := u.updateOrder.Run(request.Id_Jabon, request.Cantidad, request.Estado, request.Costo, request.Codigo_Identificador, request.Tipo); err != nil {
		log.Printf("Error actualizando la orden: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error actualizando la orden"})
		return
	}

	log.Println("Orden actualizada exitosamente")
	ctx.JSON(http.StatusOK, gin.H{"message": "Orden actualizada exitosamente"})
}