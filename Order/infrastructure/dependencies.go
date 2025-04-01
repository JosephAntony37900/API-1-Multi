package infrastructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	app_order "github.com/JosephAntony37900/API-1-Multi/Order/application"
	repo_order "github.com/JosephAntony37900/API-1-Multi/Order/infrastructure/repository"
	control_order "github.com/JosephAntony37900/API-1-Multi/Order/infrastructure/controllers"
	routes_order "github.com/JosephAntony37900/API-1-Multi/Order/infrastructure/routes"
)

func InitOrderDependencies(engine *gin.Engine, db *sql.DB){
	orderRepo := repo_order.NewOrderRepoMySQL(db)

	createOrderUseCase := app_order.NewCreateOrder(orderRepo)
	getOrderByCodigoIdUseCase := app_order.NewGetOrderByCodigoId(orderRepo)
	updateOrderUseCase := app_order.NewUpdateOrder(orderRepo)

	createOrderController := control_order.NewCreateOrderController(createOrderUseCase)
	getOrderByCodigoIdController := control_order.NewGetOrderController(getOrderByCodigoIdUseCase)
	updateOrderController := control_order.NewUpdateOrderController(updateOrderUseCase)

	routes_order.SetupOrderRoutes(engine, createOrderController, updateOrderController, getOrderByCodigoIdController )
}