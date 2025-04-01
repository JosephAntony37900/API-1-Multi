package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/API-1-Multi/Order/application"
)

type GetOrderController struct {
	getOrder *application.GetOrderByCodigoId
}

func NewGetOrderController(getOrder *application.GetOrderByCodigoId) *GetOrderController {
	return &GetOrderController{getOrder: getOrder}
}

func (g *GetOrderController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para obtener una orden por Codigo_Identificador")

	codigoIdentificador := ctx.Query("codigo_identificador")
	if codigoIdentificador == "" {
		log.Println("Falta el parámetro Codigo_Identificador")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere el parámetro 'codigo_identificador'"})
		return
	}

	log.Printf("Buscando orden con Codigo_Identificador: %s", codigoIdentificador)
	order, err := g.getOrder.Run(codigoIdentificador)
	if err != nil {
		log.Printf("Error obteniendo la orden: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo la orden"})
		return
	}

	if order == nil {
		log.Printf("No se encontró una orden con Codigo_Identificador: %s", codigoIdentificador)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Orden no encontrada"})
		return
	}

	log.Printf("Orden encontrada: %+v", order)
	ctx.JSON(http.StatusOK, gin.H{"order": order})
}