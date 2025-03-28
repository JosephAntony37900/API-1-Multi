package routes

import (
	"github.com/JosephAntony37900/API-1-Multi/Soaps/infrastructure/controllers"
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/gin-gonic/gin"
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
	// Grupo de rutas públicas ( cambiar despues a protegidas porfa)
	engine.POST("/soaps", createSoapController.Handle)
	engine.GET("/soaps", getAllSoapsController.Handle)
	engine.GET("/soaps/:id", getByIdSoapController.Handle)

	// Grupo de rutas protegidas con JWTMiddleware
	protectedRoutes := engine.Group("/soaps")
	protectedRoutes.Use(helpers.JWTMiddleware()) // Aplica el middleware

	// Rutas protegidas
	protectedRoutes.PUT("/:id", updateSoapController.Handle)
	protectedRoutes.DELETE("/:id", deleteSoapController.Handle)
	protectedRoutes.GET("/admin/:adminId", getSoapsByAdminController.Handle)
}