package infrastructure

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/application"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/controllers"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/repository"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/routes"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/rabbitmq"
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"database/sql"
)

func InitLevelDependencies(engine *gin.Engine, db *sql.DB, rabbitmqURI string) {
	if err := helpers.InitRabbitMQ(rabbitmqURI); err != nil {
		log.Fatalf("Error inicializando RabbitMQ: %v", err)
	}

	// Obtener el canal RabbitMQ
	channel := helpers.GetRabbitMQChannel()
	if channel == nil {
		log.Fatalf("RabbitMQ channel is not initialized")
	}

	levelRepo := repository.NewLevelReadingRepoMySQL(db)

	publisher := rabbitmq.NewRabbitMQPublisher("amq.topic")

	createLevelUseCase := application.NewCreateLevelReading(levelRepo)
	levelMessageService := application.NewLevelReadingMessageService(levelRepo, createLevelUseCase, publisher)

	go func() {
		err := rabbitmq.StartLevelReadingConsumer(levelMessageService, "nivel.lectura", "sensor.nivel", "amq.topic")
		if err != nil {
			log.Fatalf("Error al consumir mensajes: %v", err)
		}
	}()

	routes.SetupLevelReadingRoutes(
		engine,
		controllers.NewCreateLevelReadingController(createLevelUseCase),
		controllers.NewGetLevelReadingsController(application.NewGetLevelReading(levelRepo)),
		controllers.NewGetLevelReadingByIdController(application.NewGetByIdLevelReading(levelRepo)),
	)
}