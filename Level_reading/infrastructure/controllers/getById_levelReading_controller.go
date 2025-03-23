package controllers

import (
	"log"
	"net/http"
	"strconv"

	application "github.com/JosephAntony37900/API-1-Multi/Level_reading/application"
	"github.com/gin-gonic/gin"
)

// Controlador para obtener un nivel de lectura por ID
type GetLevelReadingByIdController struct {
	getByIdUseCase *application.GetByIdLevelReading
}

func NewGetLevelReadingByIdController(getById *application.GetByIdLevelReading) *GetLevelReadingByIdController {
	return &GetLevelReadingByIdController{getByIdUseCase: getById}
}

func (c *GetLevelReadingByIdController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para obtener un nivel de lectura por ID")

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("ID inválido: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	levelReading, err := c.getByIdUseCase.Run(id)
	if err != nil {
		log.Printf("Error obteniendo el nivel de lectura: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo el nivel de lectura"})
		return
	}

	if levelReading == nil {
		log.Println("Nivel de lectura no encontrado")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Nivel de lectura no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, levelReading)
}
