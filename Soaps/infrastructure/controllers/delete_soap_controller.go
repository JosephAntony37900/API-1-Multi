package controllers

import (
	"log"
	"strconv"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	"github.com/gin-gonic/gin"
)

type DeleteSoapController struct {
	deleteSoap *application.DeleteSoap
}

func NewDeleteSoapController(deleteSoap *application.DeleteSoap) *DeleteSoapController {
	return &DeleteSoapController{deleteSoap: deleteSoap}
}

func (c *DeleteSoapController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para eliminar un jabón")

	// Obtener el ID del jabón desde la URL
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error convirtiendo el ID a entero: %v", err)
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	log.Printf("Eliminando jabón: ID=%d", id)

	// Ejecutar el caso de uso para eliminar el jabón
	if err := c.deleteSoap.Run(id); err != nil {
		log.Printf("Error eliminando el jabón: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta de éxito
	log.Println("Jabón eliminado exitosamente")
	ctx.JSON(200, gin.H{"message": "jabón eliminado exitosamente"})
}