package infrastructure

import (
    "log"
    "database/sql"

    "github.com/gin-gonic/gin"
    app_order "github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/application"
    repo_order "github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/infrastructure/repository"
    control_order "github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/infrastructure/controllers"
    routes_order "github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/infrastructure/routes"
    rabbitmq_order "github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/infrastructure/rabbitmq"
    "github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/service"
    "github.com/JosephAntony37900/API-1-Multi/helpers"
)

func InitOrderBombaDependencies(engine *gin.Engine, db *sql.DB, rabbitmqURI string) {
    if err := helpers.InitRabbitMQ(rabbitmqURI); err != nil {
        log.Fatalf("Error inicializando RabbitMQ: %v", err)
    }

    orderRepo := repo_order.NewOrderRepoMySQL(db)
    bombaPublisher := rabbitmq_order.NewRabbitMQServoPublisher()
    orderService := service.NewOrderService(orderRepo, bombaPublisher)

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
