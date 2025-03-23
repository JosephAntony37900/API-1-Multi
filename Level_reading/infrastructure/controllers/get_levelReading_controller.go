package controllers

import (
	"log"
	"net/http"
	_"strconv"

	application "github.com/JosephAntony37900/API-1-Multi/Level_reading/application"
	"github.com/gin-gonic/gin"
)

// Controlador para obtener todos los niveles de lectura
type GetLevelReadingsController struct {
	getAllUseCase *application.GetLevelReading
}

func NewGetLevelReadingsController(getAll *application.GetLevelReading) *GetLevelReadingsController {
	return &GetLevelReadingsController{getAllUseCase: getAll}
}

func (c *GetLevelReadingsController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para obtener todos los niveles de lectura")

	levelReadings, err := c.getAllUseCase.Run()
	if err != nil {
		log.Printf("Error obteniendo los niveles de lectura: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo los niveles de lectura"})
		return
	}

	ctx.JSON(http.StatusOK, levelReadings)
}

