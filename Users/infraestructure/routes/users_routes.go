package routes

import (
    "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/controllers"
	"github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/services"
    "github.com/gin-gonic/gin"
	"os"
)

func SetupUserRoutes(r *gin.Engine, createUserController *controllers.CreateUserController, 
    loginUserController *controllers.LoginUserController, getUserController *controllers.GetUsersController, 
    deleteUserController *controllers.DeleteUserController, updateUserController *controllers.UpdateUserController) {
    
    jwtSecret := os.Getenv("JWT_SECRET")
    
    r.POST("/users", createUserController.Handle)
    r.POST("/login", loginUserController.Handle)
    
    authGroup := r.Group("/")
    authGroup.Use(service.AuthMiddleware(jwtSecret))
    {
        authGroup.GET("/users", getUserController.Handle)
        authGroup.DELETE("/users/:id", deleteUserController.Handle)
        authGroup.PUT("/users/:id", updateUserController.Handle)
    }
}
