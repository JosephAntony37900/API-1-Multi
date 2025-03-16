package main

import (
	repo_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/repository"
	"github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	"github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/controllers"
	"github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/routes"
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	// Inicializar la conexión a la base de datos
	db, err := helpers.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error inicializando la conexión a MySQL: %v", err)
	}
	defer db.Close()

	// Crear el repositorio
	soapRepo := repo_soap.NewSoapRepoMySQL(db)

	// Crear los casos de uso
	createSoapUseCase := application.NewCreateSoap(soapRepo)
	getAllSoapsUseCase := application.NewGetAllSoaps(soapRepo)
	getByIdSoapUseCase := application.NewGetByIdSoap(soapRepo)
	updateSoapUseCase := application.NewUpdateSoaps(soapRepo)
	deleteSoapUseCase := application.NewDeleteSoap(soapRepo)

	// Crear los controladores
	createSoapController := controllers.NewCreateSoapController(createSoapUseCase)
	getAllSoapsController := controllers.NewGetAllSoapsController(getAllSoapsUseCase)
	getByIdSoapController := controllers.NewGetByIdSoapController(getByIdSoapUseCase)
	updateSoapController := controllers.NewUpdateSoapController(updateSoapUseCase)
	deleteSoapController := controllers.NewDeleteSoapController(deleteSoapUseCase)

	// Configurar el enrutador
	engine := gin.Default()

	// Configurar las rutas
	routes.SetupRoutes(
		engine,
		createSoapController,
		getAllSoapsController,
		getByIdSoapController,
		updateSoapController,
		deleteSoapController,
	)

	// Iniciar el servidor
	engine.Run(":8000")
}