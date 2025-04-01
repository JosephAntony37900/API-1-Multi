package infrastructure

import (
	"github.com/gin-gonic/gin"
	app_users "github.com/JosephAntony37900/API-1-Multi/Users/application"
	control_users "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/controllers"
	repo_users "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/repository"
	routes_users "github.com/JosephAntony37900/API-1-Multi/Users/infraestructure/routes"
	"database/sql"
)

func InitUserDependencies(engine *gin.Engine, db *sql.DB) {
	userRepo := repo_users.NewCreateUserRepoMySQL(db)

	createUsersUseCase := app_users.NewCreateUser(userRepo)
	getAllUsersUseCase := app_users.NewGetUsers(userRepo)
	deleteUsersUseCase := app_users.NewDeleteUser(userRepo)
	loginUserUseCase := app_users.NewLoginUser(userRepo)
	updateUsersUseCase := app_users.NewUpdateUser(userRepo)

	createUserController := control_users.NewCreateUserController(createUsersUseCase)
	getAllUsersController := control_users.NewUsersController(getAllUsersUseCase)
	deleteUsersController := control_users.NewDeleteUserController(deleteUsersUseCase)
	loginUserController := control_users.NewLoginUserController(loginUserUseCase)
	updateUsersController := control_users.NewUpdateUserController(updateUsersUseCase)

	routes_users.SetupUserRoutes(engine, createUserController, loginUserController, getAllUsersController, deleteUsersController, updateUsersController)
}