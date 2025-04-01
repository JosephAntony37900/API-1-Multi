package infrastructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	app_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	control_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/controllers"
	repo_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/repository"
	routes_soap "github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/routes"
)

func InitSoapDependencies(engine *gin.Engine, db *sql.DB) {

	soapRepo := repo_soap.NewSoapRepoMySQL(db)

	createSoapUseCase := app_soap.NewCreateSoap(soapRepo)
	getAllSoapsUseCase := app_soap.NewGetAllSoaps(soapRepo)
	getByIdSoapUseCase := app_soap.NewGetByIdSoap(soapRepo)
	updateSoapUseCase := app_soap.NewUpdateSoaps(soapRepo)
	deleteSoapUseCase := app_soap.NewDeleteSoap(soapRepo)
	getByAdminUseCase := app_soap.NewGetSoapsByAdmin(soapRepo)

	createSoapController := control_soap.NewCreateSoapController(createSoapUseCase)
	getAllSoapsController := control_soap.NewGetAllSoapsController(getAllSoapsUseCase)
	getByIdSoapController := control_soap.NewGetByIdSoapController(getByIdSoapUseCase)
	updateSoapController := control_soap.NewUpdateSoapController(updateSoapUseCase)
	deleteSoapController := control_soap.NewDeleteSoapController(deleteSoapUseCase)
	getByAdminController := control_soap.NewGetSoapsByAdminController(getByAdminUseCase)

	routes_soap.SetupRoutes(
		engine,
		createSoapController,
		getAllSoapsController,
		getByIdSoapController,
		updateSoapController,
		deleteSoapController,
		getByAdminController,
	)

}