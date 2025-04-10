package controllers

import (
	"log"
	"net/http"

	"github.com/JosephAntony37900/API-1-Multi/Order/application"
	"github.com/JosephAntony37900/API-1-Multi/Order/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Order/domain/service"
	"github.com/gin-gonic/gin"
)

type UpdateOrderController struct {
	updateOrder  *application.UpdateOrder
	orderService *service.OrderService
}

func NewUpdateOrderController(updateOrder *application.UpdateOrder, orderService *service.OrderService) *UpdateOrderController {
	return &UpdateOrderController{
		updateOrder:  updateOrder,
		orderService: orderService,
	}
}

func (u *UpdateOrderController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para actualizar una orden")

	var request struct {
		Id_Jabon             int     `json:"id_jabon"`
		Cantidad             float64 `json:"cantidad"`
		Estado               int     `json:"estado"`
		Costo                float64 `json:"costo"`
		Codigo_Identificador string  `json:"codigo_identificador"`
		Tipo                 bool    `json:"tipo"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cuerpo de la solicitud inválido"})
		return
	}

	order := entities.Order{
		Id_Jabon:             request.Id_Jabon,
		Cantidad:             request.Cantidad,
		Estado:               request.Estado,
		Costo:                request.Costo,
		Codigo_Identificador: request.Codigo_Identificador,
		Tipo:                 request.Tipo,
	}

	if err := u.updateOrder.Run(order.Id_Jabon, order.Cantidad, order.Estado, order.Costo, order.Codigo_Identificador, order.Tipo); err != nil {
		log.Printf("Error actualizando la orden: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error actualizando la orden"})
		return
	}

	log.Println("Orden actualizada exitosamente")

	if !order.Tipo && order.Estado == 2 {
		tiempoDespacho := int(order.Cantidad * 5)
		err := u.orderService.ProcessOrder(
			order.Codigo_Identificador,
			tiempoDespacho,
			order.Tipo,
		)
		if err != nil {
			log.Printf("Error procesando el despacho: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error procesando el despacho"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Orden procesada exitosamente"})
}
