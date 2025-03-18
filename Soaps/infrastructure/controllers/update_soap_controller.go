package controllers

import (
	"log"
	"strconv"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	"github.com/gin-gonic/gin"
)

type UpdateSoapController struct {
	updateSoap *application.UpdateSoap
}

func NewUpdateSoapController(updateSoap *application.UpdateSoap) *UpdateSoapController {
	return &UpdateSoapController{updateSoap: updateSoap}
}

func (c *UpdateSoapController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para actualizar un jabón")

	// Obtener el ID del jabón desde la URL
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error convirtiendo el ID a entero: %v", err)
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	// Estructura para decodificar el JSON de la solicitud
	var request struct {
		Nombre   string  `json:"nombre"`
		Marca    string  `json:"marca"`
		Tipo     string  `json:"tipo"`
		Precio   float64 `json:"precio"`
		Densidad float64 `json:"densidad"`
	}

	// Decodificar el cuerpo de la solicitud
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando el cuerpo de la solicitud: %v", err)
		ctx.JSON(400, gin.H{"error": "cuerpo de la solicitud inválido"})
		return
	}

	log.Printf("Actualizando jabón: ID=%d, Nombre=%s, Marca=%s, Tipo=%s, Precio=%f, Densidad=%f",
		id, request.Nombre, request.Marca, request.Tipo, request.Precio, request.Densidad)

	// Ejecutar el caso de uso para actualizar el jabón
	if err := c.updateSoap.Run(id, request.Nombre, request.Marca, request.Tipo, request.Precio, request.Densidad); err != nil {
		log.Printf("Error actualizando el jabón: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta de éxito
	log.Println("Jabón actualizado exitosamente")
	ctx.JSON(200, gin.H{"message": "jabón actualizado exitosamente"})
}