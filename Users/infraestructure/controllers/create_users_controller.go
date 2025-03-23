package controllers

import (
	"log"

	"github.com/JosephAntony37900/API-1-Multi/Users/application"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	CreateUsers *application.CreateUsers
}

func NewCreateUserController(CreateUsers *application.CreateUsers) *CreateUserController {
	return &CreateUserController{CreateUsers: CreateUsers}
}

func (c *CreateUserController) Handle(ctx *gin.Context) {
	log.Println("Petición de crear un producto, recibido")

	var request struct {
		Nombre     string `json:"Nombre"`
		Email      string `json:"Email"`
		Contraseña string `json:"Contrasena"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		log.Printf("Error decodificando la petición del body: %v", err)
		ctx.JSON(400, gin.H{"error": "petición del body invalida"})
		return
	}
	log.Printf("Creando usuario: Nombre=%s, email=%s, contraseña=%s", request.Nombre, request.Email, request.Contraseña)

	if err := c.CreateUsers.Run(request.Nombre, request.Email, request.Contraseña); err != nil {
		log.Printf("Error creando el usuario: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Usuario creado exitosamente")
	ctx.JSON(201, gin.H{"message": "usuario creado exitosamente"})

}
