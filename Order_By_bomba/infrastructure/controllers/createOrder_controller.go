package controllers

import (
	"log"
	"net/http"

	"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/application"
	_"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/service"
	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
    createOrUpdateOrder *application.CreateOrder
    orderService        *service.OrderService
}

func NewCreateOrderController(createOrUpdateOrder *application.CreateOrder, orderService *service.OrderService) *CreateOrderController {
    return &CreateOrderController{
        createOrUpdateOrder: createOrUpdateOrder,
        orderService:        orderService,
    }
}

func (c *CreateOrderController) Handle(ctx *gin.Context) {
    log.Println("Recibe la petición para crear o actualizar una orden")

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

    var tiempoDespacho int
    switch int(request.Cantidad) {
    case 1:
        tiempoDespacho = 5
    case 2:
        tiempoDespacho = 10
    case 3:
        tiempoDespacho = 15
    default:
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cantidad inválida. Use 1, 2 o 3"})
        return
    }

    err := c.orderService.ProcessOrder(
        request.Codigo_Identificador,
        tiempoDespacho,
        "",
        request.Tipo,
    )
    if err != nil {
        log.Printf("Error procesando orden: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    err = c.createOrUpdateOrder.Run(request.Cantidad, request.Estado, request.Costo, request.Codigo_Identificador, request.Tipo)
    if err != nil {
        log.Printf("Error creando/actualizando la orden: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando/actualizando la orden"})
        return
    }

    log.Println("Orden creada y proceso de despacho iniciado exitosamente")
    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Orden procesada exitosamente",
        "tiempo_despacho": tiempoDespacho,
    })
}