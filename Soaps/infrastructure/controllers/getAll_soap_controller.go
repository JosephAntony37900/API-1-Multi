package controllers

import (
	"log"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	"github.com/gin-gonic/gin"
)

type GetAllSoapsController struct {
	getAllSoaps *application.GetAllSoaps
}

func NewGetAllSoapsController(getAllSoaps *application.GetAllSoaps) *GetAllSoapsController {
	return &GetAllSoapsController{getAllSoaps: getAllSoaps}
}

func (c *GetAllSoapsController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para obtener todos los jabones")

	// Ejecutar el caso de uso para obtener todos los jabones
	soaps, err := c.getAllSoaps.Run()
	if err != nil {
		log.Printf("Error obteniendo los jabones: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Respuesta con la lista de jabones
	log.Println("Jabones obtenidos exitosamente")
	ctx.JSON(200, soaps)
}