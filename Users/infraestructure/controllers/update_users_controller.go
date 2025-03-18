package controllers

import (
    "github.com/JosephAntony37900/API-1-Multi/Users/application"
    "github.com/gin-gonic/gin"
    "strconv"
)

type UpdateUserController struct {
	updateUser *application.UpdateUser
}

func NewUpdateUserController(updateUser *application.UpdateUser) *UpdateUserController{
	return &UpdateUserController{updateUser: updateUser}
}

func (update *UpdateUserController) Handle(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil { 
		ctx.JSON(400, gin.H{"error": "ID de usuario invalido"})
		return
	}

	var request struct {
		Nombre string `json:"Nombre"`
		Email string `json:"Email"`
		Contraseña string `json:"Contraseña"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "petición del body invalida"})
		return
	}

	if err := update.updateUser.Run(id, request.Email, request.Nombre, request.Contraseña); err != nil{
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "usuario actualizado correctamente"})
}