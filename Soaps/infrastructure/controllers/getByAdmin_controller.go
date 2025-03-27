package controllers

import (
	"log"
	"strconv"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/application"
	"github.com/gin-gonic/gin"
)

type GetSoapsByAdminController struct {
	getSoapsByAdmin *application.GetSoapsByAdmin
}

func NewGetSoapsByAdminController(getSoapsByAdmin *application.GetSoapsByAdmin) *GetSoapsByAdminController {
	return &GetSoapsByAdminController{getSoapsByAdmin: getSoapsByAdmin}
}

func (c *GetSoapsByAdminController) Handle(ctx *gin.Context) {
	log.Println("Recibe la petición para obtener jabones de un administrador")

	adminIdStr := ctx.Param("adminId")
	adminId, err := strconv.Atoi(adminIdStr)
	if err != nil {
		log.Printf("Error convirtiendo el ID del administrador a entero: %v", err)
		ctx.JSON(400, gin.H{"error": "ID del administrador inválido"})
		return
	}

	log.Printf("Obteniendo jabones del administrador: ID=%d", adminId)

	soaps, err := c.getSoapsByAdmin.Run(adminId)
	if err != nil {
		log.Printf("Error obteniendo jabones: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Jabones obtenidos exitosamente")
	ctx.JSON(200, soaps)
}