package routes

import (
    "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, createUserController *controllers.CreateUserController, loginUserController *controllers.LoginUserController, getUserController *controllers.GetUsersController, deleteUserController *controllers.DeleteUserController, updateUserController *controllers.UpdateUserController) {
	r.POST("/users", createUserController.Handle)
	r.POST("/login", loginUserController.Handle) 
	r.GET("/users", getUserController.Handle)
	r.DELETE("/users/:id", deleteUserController.Handle)
	r.PUT("/users/:id", updateUserController.Handle)
}
