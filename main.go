package main

import (
	repo_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/repository"
	app_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	control_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/controllers"
	routes_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/routes"
	repo_users "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/repository"
	app_users "github.com/JosephAntony37900/API-1-Multi/Users/application"
	control_users "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/controllers"
	routes_users "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/routes"
	repo_levels "github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/repository"
	app_levels "github.com/JosephAntony37900/API-1-Multi/Level_reading/application"
	control_levels "github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/controllers"
	routes_levels "github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/routes"
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

	// repositorios
	soapRepo := repo_soap.NewSoapRepoMySQL(db)
	userRepo := repo_users.NewCreateUserRepoMySQL(db)
	levelRepo := repo_levels.NewLevelReadingRepoMySQL(db)


	// casos de uso de soap
	createSoapUseCase := app_soap.NewCreateSoap(soapRepo)
	getAllSoapsUseCase := app_soap.NewGetAllSoaps(soapRepo)
	getByIdSoapUseCase := app_soap.NewGetByIdSoap(soapRepo)
	updateSoapUseCase := app_soap.NewUpdateSoaps(soapRepo)
	deleteSoapUseCase := app_soap.NewDeleteSoap(soapRepo)

	// casos de uso de users
	createUsersUseCase := app_users.NewCreateUser(userRepo)
	getAllUsersUseCase := app_users.NewGetUsers(userRepo)
	deleteUsersUseCase := app_users.NewDeleteUser(userRepo)
	loginUserUseCase := app_users.NewLoginUser(userRepo)
	updateUsersUseCase := app_users.NewUpdateUser(userRepo)

	// casos de uso de Level_reading
	createLevelUseCase := app_levels.NewCreateLevelReading(levelRepo)
	getAllLevelsUseCase := app_levels.NewGetLevelReading(levelRepo)
	getByIdLevelUseCase := app_levels.NewGetByIdLevelReading(levelRepo)
	
	// Crontoladores de soap
	createSoapController := control_soap.NewCreateSoapController(createSoapUseCase)
	getAllSoapsController := control_soap.NewGetAllSoapsController(getAllSoapsUseCase)
	getByIdSoapController := control_soap.NewGetByIdSoapController(getByIdSoapUseCase)
	updateSoapController := control_soap.NewUpdateSoapController(updateSoapUseCase)
	deleteSoapController := control_soap.NewDeleteSoapController(deleteSoapUseCase)

	// controladores de users
	createUserController := control_users.NewCreateUserController(createUsersUseCase)
	getAllUsersController := control_users.NewUsersController(getAllUsersUseCase)
	deleteUsersController := control_users.NewDeleteUserController(deleteUsersUseCase)
	loginUserController := control_users.NewLoginUserController(loginUserUseCase)
	updateUsersController := control_users.NewUpdateUserController(updateUsersUseCase)

	// Controladores de Level_reading
	createLevelController := control_levels.NewCreateLevelReadingController(createLevelUseCase)
	getAllLevelsController := control_levels.NewGetLevelReadingsController(getAllLevelsUseCase)
	getByIdLevelController := control_levels.NewGetLevelReadingByIdController(getByIdLevelUseCase)
	
	// Configurar el enrutador
	engine := gin.Default()

	// Configurar las rutas de soap
	routes_soap.SetupRoutes(
		engine,
		createSoapController,
		getAllSoapsController,
		getByIdSoapController,
		updateSoapController,
		deleteSoapController,
	)

	routes_users.SetupUserRoutes(engine, createUserController, loginUserController, getAllUsersController, deleteUsersController, updateUsersController)

    // Configurar las rutas de Level_reading
	routes_levels.SetupLevelReadingRoutes(
		engine,
		createLevelController,
		getAllLevelsController,
		getByIdLevelController,
		
	)

	// Iniciar el servidor
	engine.Run(":8000")
	
}