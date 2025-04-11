package infrastructure

import (
	"database/sql"
	"log"

	app_order "github.com/JosephAntony37900/API-1-Multi/Order/application"
	"github.com/JosephAntony37900/API-1-Multi/Order/domain/service"
	control_order "github.com/JosephAntony37900/API-1-Multi/Order/infrastructure/controllers"
	rabbitmq_order "github.com/JosephAntony37900/API-1-Multi/Order/infrastructure/rabbitmq"
	repo_order "github.com/JosephAntony37900/API-1-Multi/Order/infrastructure/repository"
	routes_order "github.com/JosephAntony37900/API-1-Multi/Order/infrastructure/routes"
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/gin-gonic/gin"
)

func InitOrderDependencies(engine *gin.Engine, db *sql.DB, rabbitmqURI string) {
	if err := helpers.InitRabbitMQ(rabbitmqURI); err != nil {
		log.Fatalf("Error inicializando RabbitMQ: %v", err)
	}

	orderRepo := repo_order.NewOrderRepoMySQL(db)
	bombaPublisher := rabbitmq_order.NewRabbitMQBombaPublisher()
	servoPublisher := rabbitmq_order.NewRabbitMQServoPublisher()
	orderService := service.NewOrderService(orderRepo, servoPublisher, bombaPublisher)

	createOrderUseCase := app_order.NewCreateOrder(orderRepo)
	getOrderByCodigoIdUseCase := app_order.NewGetOrderByCodigoId(orderRepo)
	updateOrderUseCase := app_order.NewUpdateOrder(orderRepo)

	createOrderController := control_order.NewCreateOrderController(createOrderUseCase, orderService)
	getOrderByCodigoIdController := control_order.NewGetOrderController(getOrderByCodigoIdUseCase)
	updateOrderController := control_order.NewUpdateOrderController(updateOrderUseCase, orderService)

	routes_order.SetupOrderRoutes(engine, createOrderController, updateOrderController, getOrderByCodigoIdController)

}
