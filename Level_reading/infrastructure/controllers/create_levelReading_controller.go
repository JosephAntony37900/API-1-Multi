package controllers

import (
	"log"
	"net/http"
	_"strconv"

	application "github.com/JosephAntony37900/API-1-Multi/Level_reading/application"
	"github.com/gin-gonic/gin"
)

type CreateLevelReadingController struct {
	createUseCase *application.CreateLevelReading
}

func NewCreateLevelReadingController(create *application.CreateLevelReading) *CreateLevelReadingController {
	return &CreateLevelReadingController{createUseCase: create}
}

func (c *CreateLevelReadingController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para crear un nivel de lectura")

	var request struct {
		Fecha       int     `json:"fecha"`
		Id_Jabon    int     `json:"id_jabon"`
		Nivel_Jabon float64 `json:"nivel_jabon"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cuerpo de la solicitud inválido"})
		return
	}

	log.Printf("Creando nivel de lectura: Fecha=%d, Id_Jabon=%d, Nivel_Jabon=%f",
		request.Fecha, request.Id_Jabon, request.Nivel_Jabon)

	if err := c.createUseCase.Run(request.Fecha, request.Id_Jabon, request.Nivel_Jabon); err != nil {
		log.Printf("Error creando el nivel de lectura: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Nivel de lectura creado exitosamente")
	ctx.JSON(http.StatusCreated, gin.H{"message": "Nivel de lectura creado exitosamente"})
}