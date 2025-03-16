package routes

import (
	"github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	engine *gin.Engine,
	createSoapController *controllers.CreateSoapController,
	getAllSoapsController *controllers.GetAllSoapsController,
	getByIdSoapController *controllers.GetByIdSoapController,
	updateSoapController *controllers.UpdateSoapController,
	deleteSoapController *controllers.DeleteSoapController,
) {
	// Ruta para crear un jabón
	engine.POST("/soaps", createSoapController.Handle)

	// Ruta para obtener todos los jabones
	engine.GET("/soaps", getAllSoapsController.Handle)

	// Ruta para obtener un jabón por ID
	engine.GET("/soaps/:id", getByIdSoapController.Handle)

	// Ruta para actualizar un jabón por ID
	engine.PUT("/soaps/:id", updateSoapController.Handle)

	// Ruta para eliminar un jabón por ID
	engine.DELETE("/soaps/:id", deleteSoapController.Handle)
}