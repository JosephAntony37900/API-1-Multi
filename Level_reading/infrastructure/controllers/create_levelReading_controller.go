package controllers

import (
	"log"
	"net/http"
	"time"

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
		Fecha       string  `json:"fecha"`       // Fecha como string (ISO8601: "2023-03-22T15:04:05Z")
		Id_Jabon    int     `json:"id_jabon"`   
		Nivel_Jabon float64 `json:"nivel_jabon"` 
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cuerpo de la solicitud inválido"})
		return
	}

	fecha, err := time.Parse(time.RFC3339, request.Fecha) //le doy el formato ISO8601
	if err != nil {
		log.Printf("Error convirtiendo la fecha: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha inválido"})
		return
	}

	// Convierto time.Time a timestamp (int64)
	timestamp := fecha.Unix()

	log.Printf("Creando nivel de lectura: Fecha=%s, Id_Jabon=%d, Nivel_Jabon=%f",
		fecha.String(), request.Id_Jabon, request.Nivel_Jabon)

	if err := c.createUseCase.Run(timestamp, request.Id_Jabon, request.Nivel_Jabon); err != nil {
		log.Printf("Error creando el nivel de lectura: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Nivel de lectura creado exitosamente")
	ctx.JSON(http.StatusCreated, gin.H{"message": "Nivel de lectura creado exitosamente"})
}