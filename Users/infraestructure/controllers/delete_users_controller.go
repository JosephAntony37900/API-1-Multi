package controllers

import (
    "github.com/JosephAntony37900/API-1-Multi/Users/application"
    "github.com/gin-gonic/gin"
    "strconv"
)

type DeleteUserController struct {
	deleteUser *application.DeleteUser
}

func NewDeleteUserController(deleteUser *application.DeleteUser) *DeleteUserController{
	return &DeleteUserController{deleteUser: deleteUser}
}

func (du *DeleteUserController) Handle(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de usuario invalido"})
		return
	}

	if err := du.deleteUser.Run(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "usuario eliminado correctamente"})
}