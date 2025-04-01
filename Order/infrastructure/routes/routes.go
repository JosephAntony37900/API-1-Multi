package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/API-1-Multi/Order/infrastructure/controllers"
)

func SetupOrderRoutes(router *gin.Engine, createController *controllers.CreateOrderController, updateController *controllers.UpdateOrderController, getController *controllers.GetOrderController) {
	router.POST("/orders", createController.Handle)      
	router.PUT("/orders", updateController.Handle)      
	router.GET("/orders", getController.Handle)         
}