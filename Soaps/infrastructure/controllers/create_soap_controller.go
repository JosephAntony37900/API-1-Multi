package controllers

import (
	"log"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	"github.com/gin-gonic/gin"
)

type CreateSoapController struct {
	createSoap *application.CreateSoap
}

func NewCreateSoapController(createSoap *application.CreateSoap) *CreateSoapController {
	return &CreateSoapController{createSoap: createSoap}
}

func (c *CreateSoapController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para crear un jabón")

	// Estructura para decodificar el JSON de la solicitud
	var request struct {
		Nombre   string  `json:"nombre"`
		Marca    string  `json:"marca"`
		Tipo     string  `json:"tipo"`
		Precio   float64 `json:"precio"`
		Densidad float64 `json:"densidad"`
		Id_Usuario_Admin int `json:"id_usuario_admin "`
	}

	// Decodificar el cuerpo de la solicitud
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando el cuerpo de la solicitud: %v", err)
		ctx.JSON(400, gin.H{"error": "cuerpo de la solicitud inválido"})
		return
	}

	log.Printf("Creando jabón: Nombre=%s, Marca=%s, Tipo=%s, Precio=%f, Densidad=%f",
		request.Nombre, request.Marca, request.Tipo, request.Precio, request.Densidad)

	// Ejecutar el caso de uso para crear el jabón
	if err := c.createSoap.Run(request.Nombre, request.Marca, request.Tipo, request.Precio, request.Densidad, request.Id_Usuario_Admin); err != nil {
		log.Printf("Error creando el jabón: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta de éxito
	log.Println("Jabón creado exitosamente")
	ctx.JSON(201, gin.H{"message": "jabón creado exitosamente"})
}