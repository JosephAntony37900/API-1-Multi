package routes

import (
	"github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/controllers"
	"github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/services"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupRoutes(
	engine *gin.Engine,
	createSoapController *controllers.CreateSoapController,
	getAllSoapsController *controllers.GetAllSoapsController,
	getByIdSoapController *controllers.GetByIdSoapController,
	updateSoapController *controllers.UpdateSoapController,
	deleteSoapController *controllers.DeleteSoapController,
	getSoapsByAdminController *controllers.GetSoapsByAdminController,
) {

	jwtSecret := os.Getenv("JWT_SECRET")

	// Grupo de rutas públicas ( cambiar despues a protegidas porfa)
	engine.POST("/soaps", createSoapController.Handle)
	engine.GET("/soaps", getAllSoapsController.Handle)
	engine.GET("/soaps/:id", getByIdSoapController.Handle)

	protectedRoutes := engine.Group("/soaps")
	protectedRoutes.Use(service.AuthMiddleware(jwtSecret))

	protectedRoutes.PUT("/:id", updateSoapController.Handle)
	protectedRoutes.DELETE("/:id", deleteSoapController.Handle)
	protectedRoutes.GET("/admin/:adminId", getSoapsByAdminController.Handle)
}