package infrastructure

import (
    "log"
    "database/sql"

    "github.com/gin-gonic/gin"
    app_order "github.com/JosephAntony37900/API-1-Multi/Order_By_servo/application"
    repo_order "github.com/JosephAntony37900/API-1-Multi/Order_By_servo/infrastructure/repository"
    control_order "github.com/JosephAntony37900/API-1-Multi/Order_By_servo/infrastructure/controllers"
    routes_order "github.com/JosephAntony37900/API-1-Multi/Order_By_servo/infrastructure/routes"
    rabbitmq_order "github.com/JosephAntony37900/API-1-Multi/Order_By_servo/infrastructure/rabbitmq"
    "github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/service"
    "github.com/JosephAntony37900/API-1-Multi/helpers"
)

func InitOrderDependencies(engine *gin.Engine, db *sql.DB, rabbitmqURI string) {
    if err := helpers.InitRabbitMQ(rabbitmqURI); err != nil {
        log.Fatalf("Error inicializando RabbitMQ: %v", err)
    }

    orderRepo := repo_order.NewOrderRepoMySQL(db)
    servoPublisher := rabbitmq_order.NewRabbitMQServoPublisher()
    orderService := service.NewOrderService(orderRepo, servoPublisher)

    createOrderUseCase := app_order.NewCreateOrder(orderRepo)
    getOrderByCodigoIdUseCase := app_order.NewGetOrderByCodigoId(orderRepo)
    updateOrderUseCase := app_order.NewUpdateOrder(orderRepo)

    createOrderController := control_order.NewCreateOrderController(createOrderUseCase, orderService)
    getOrderByCodigoIdController := control_order.NewGetOrderController(getOrderByCodigoIdUseCase)
    updateOrderController := control_order.NewUpdateOrderController(updateOrderUseCase, orderService)

    routes_order.SetupOrderRoutes(engine, createOrderController, updateOrderController, getOrderByCodigoIdController)

    go func() {
        err := rabbitmq_order.StartInfraredConsumer(orderService)
        if err != nil {
            log.Fatalf("Error al consumir mensajes de RabbitMQ para Order: %v", err)
        }
    }()
}
