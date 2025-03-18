package controllers

import (
	"log"
	"strconv"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	"github.com/gin-gonic/gin"
)

type GetByIdSoapController struct {
	getByIdSoap *application.GetByIdSoap
}

func NewGetByIdSoapController(getByIdSoap *application.GetByIdSoap) *GetByIdSoapController {
	return &GetByIdSoapController{getByIdSoap: getByIdSoap}
}

func (c *GetByIdSoapController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para obtener un jabón por ID")

	// Obtener el ID del jabón desde la URL
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error convirtiendo el ID a entero: %v", err)
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	log.Printf("Obteniendo jabón: ID=%d", id)

	// Ejecutar el caso de uso para obtener el jabón
	soap, err := c.getByIdSoap.Run(id)
	if err != nil {
		log.Printf("Error obteniendo el jabón: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta con el jabón encontrado
	log.Println("Jabón obtenido exitosamente")
	ctx.JSON(200, soap)
}