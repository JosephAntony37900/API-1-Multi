package controllers

import (
	"log"
	"net/http"
	_"strconv"

	"github.com/JosephAntony37900/API-1-Multi/Users/application"
	"github.com/gin-gonic/gin"
)

type CreateClientController struct {
	CreateClients *application.CreateClients
}

func NewCreateClientController(createClients *application.CreateClients) *CreateClientController {
	return &CreateClientController{CreateClients: createClients}
}

func (c *CreateClientController) Handle(ctx *gin.Context) {
	log.Println("Petición para crear un usuario cliente, recibido")

	var request struct {
		Nombre              string `json:"Nombre"`
		Email               string `json:"Email"`
		Contraseña          string `json:"Contrasena"`
		Codigo_Identificador string `json:"Codigo_Identificador"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando la petición del body: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "petición del body inválida"})
		return
	}

	log.Printf("Creando usuario cliente: Nombre=%s, Email=%s", request.Nombre, request.Email)

	err := c.CreateClients.Run(request.Nombre, request.Email, request.Contraseña, request.Codigo_Identificador)
	if err != nil {
		log.Printf("Error creando el usuario cliente: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Usuario cliente creado exitosamente")
	ctx.JSON(http.StatusCreated, gin.H{"message": "Usuario cliente creado exitosamente"})
}