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

	// Rutas públicas (no requieren autenticación)
	engine.GET("/soaps", getAllSoapsController.Handle)
	engine.GET("/soaps/:id", getByIdSoapController.Handle)

	// Rutas protegidas (requieren token JWT)
	protectedRoutes := engine.Group("/soaps")
	protectedRoutes.Use(service.AuthMiddleware(jwtSecret))

	protectedRoutes.POST("", createSoapController.Handle) // ← ahora protegida
	protectedRoutes.PUT("/:id", updateSoapController.Handle)
	protectedRoutes.DELETE("/:id", deleteSoapController.Handle)
	protectedRoutes.GET("/admin/:adminId", getSoapsByAdminController.Handle)
}
