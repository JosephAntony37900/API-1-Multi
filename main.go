package main

import (
	"log"
	"fmt"
	"os"

	app_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	control_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/controllers"
	repo_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/repository"
	routes_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/routes"
	app_users "github.com/JosephAntony37900/API-1-Multi/Users/application"
	control_users "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/controllers"
	repo_users "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/repository"
	routes_users "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/routes"
	repo_levels "github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/repository"
	app_levels "github.com/JosephAntony37900/API-1-Multi/Level_reading/application"
	control_levels "github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/controllers"
	routes_levels "github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/routes"
	rabbitmq "github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/rabbitmq"
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

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

	if err := helpers.InitRabbitMQ(rabbitmqURI); err != nil {
		log.Fatalf("Error inicializando RabbitMQ: %v", err)
	}
	defer helpers.CloseRabbitMQ()

	// Repos
	soapRepo := repo_soap.NewSoapRepoMySQL(db)
	userRepo := repo_users.NewCreateUserRepoMySQL(db)
	levelRepo := repo_levels.NewLevelReadingRepoMySQL(db)

	publisher := rabbitmq.NewRabbitMQPublisher("amq.topic") 

	createLevelUseCase := app_levels.NewCreateLevelReading(levelRepo)
	levelMessageService := app_levels.NewLevelReadingMessageService(levelRepo, createLevelUseCase, publisher)

	// Iniciar consumo de mensajes desde RabbitMQ
	go func() {
		err = rabbitmq.StartLevelReadingConsumer(levelMessageService, "nivel.lectura", "sensor.nivel", "amp.topic")
		if err != nil {
			log.Fatalf("Error al consumir mensajes: %v", err)
		}
	}()

	// Casos de uso de soap
	createSoapUseCase := app_soap.NewCreateSoap(soapRepo)
	getAllSoapsUseCase := app_soap.NewGetAllSoaps(soapRepo)
	getByIdSoapUseCase := app_soap.NewGetByIdSoap(soapRepo)
	updateSoapUseCase := app_soap.NewUpdateSoaps(soapRepo)
	deleteSoapUseCase := app_soap.NewDeleteSoap(soapRepo)

	// Casos de uso de users
	createUsersUseCase := app_users.NewCreateUser(userRepo)
	getAllUsersUseCase := app_users.NewGetUsers(userRepo)
	deleteUsersUseCase := app_users.NewDeleteUser(userRepo)
	loginUserUseCase := app_users.NewLoginUser(userRepo)
	updateUsersUseCase := app_users.NewUpdateUser(userRepo)

	// Controladores de soap
	createSoapController := control_soap.NewCreateSoapController(createSoapUseCase)
	getAllSoapsController := control_soap.NewGetAllSoapsController(getAllSoapsUseCase)
	getByIdSoapController := control_soap.NewGetByIdSoapController(getByIdSoapUseCase)
	updateSoapController := control_soap.NewUpdateSoapController(updateSoapUseCase)
	deleteSoapController := control_soap.NewDeleteSoapController(deleteSoapUseCase)

	// Controladores de users
	createUserController := control_users.NewCreateUserController(createUsersUseCase)
	getAllUsersController := control_users.NewUsersController(getAllUsersUseCase)
	deleteUsersController := control_users.NewDeleteUserController(deleteUsersUseCase)
	loginUserController := control_users.NewLoginUserController(loginUserUseCase)
	updateUsersController := control_users.NewUpdateUserController(updateUsersUseCase)

	// Configurar el enrutador
	engine := gin.Default()

	engine.Use(helpers.SetupCORS())

	// Configurar rutas de soap
	routes_soap.SetupRoutes(
		engine,
		createSoapController,
		getAllSoapsController,
		getByIdSoapController,
		updateSoapController,
		deleteSoapController,
	)

	routes_users.SetupUserRoutes(engine, createUserController, loginUserController, getAllUsersController, deleteUsersController, updateUsersController)

	// Configurar rutas de Level_reading
	routes_levels.SetupLevelReadingRoutes(
		engine,
		control_levels.NewCreateLevelReadingController(createLevelUseCase),
		control_levels.NewGetLevelReadingsController(app_levels.NewGetLevelReading(levelRepo)),
		control_levels.NewGetLevelReadingByIdController(app_levels.NewGetByIdLevelReading(levelRepo)),
	)

	engine.Run(":8000")
}