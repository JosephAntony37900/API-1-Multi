package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_servo/infrastructure/controllers"
)

func SetupOrderRoutes(router *gin.Engine, createController *controllers.CreateOrderController, updateController *controllers.UpdateOrderController, getController *controllers.GetOrderController) {
	router.POST("/order", createController.Handle)      
	router.PUT("/order", updateController.Handle)      
	router.GET("/order", getController.Handle)         
}