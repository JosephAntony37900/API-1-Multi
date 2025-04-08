package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/infrastructure/controllers"
)

func SetupOrderRoutes(router *gin.Engine, createController *controllers.CreateOrderController, updateController *controllers.UpdateOrderController, getController *controllers.GetOrderController) {
	router.POST("/order/bomba", createController.Handle)      
	router.PUT("/order/bomba", updateController.Handle)      
	router.GET("/order/bomba", getController.Handle)         
}