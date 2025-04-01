package main

import (
	"log"
	"fmt"
	"os"

	soap_infra "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure"
	user_infra "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure"
	level_infra "github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure"
	order_infra "github.com/JosephAntony37900/API-1-Multi/Order/infrastructure"
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	// Configuración de RabbitMQ
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPassword, rabbitmqHost, rabbitmqPort)

	db, err := helpers.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error inicializando la conexión a MySQL: %v", err)
	}
	defer db.Close()

	engine := gin.Default()
	engine.Use(helpers.SetupCORS())

	user_infra.InitUserDependencies(engine, db)
	soap_infra.InitSoapDependencies(engine, db)
	order_infra.InitOrderDependencies(engine, db)
	level_infra.InitLevelDependencies(engine, db, rabbitmqURI)

	engine.Run(":8000")
}