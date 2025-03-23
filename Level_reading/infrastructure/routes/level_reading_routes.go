package routes

import (
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupLevelReadingRoutes(
	engine *gin.Engine,
	createLevelReadingController *controllers.CreateLevelReadingController,
	getAllLevelReadingsController *controllers.GetLevelReadingsController,
	getByIdLevelReadingController *controllers.GetLevelReadingByIdController,
) {
	// Ruta para crear un nivel de lectura
	engine.POST("/level_readings", createLevelReadingController.Handle)

	// Ruta para obtener todos los niveles de lectura
	engine.GET("/level_readings", getAllLevelReadingsController.Handle)

	// Ruta para obtener un nivel de lectura por ID
	engine.GET("/level_readings/:id", getByIdLevelReadingController.Handle)
}