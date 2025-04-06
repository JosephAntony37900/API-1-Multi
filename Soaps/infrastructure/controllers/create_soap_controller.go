package controllers

import (
	"fmt"
	"log"
	"strconv"
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
	adminID, exists := ctx.Get("userID")
	if !exists {
		log.Println("No se encontró el ID de usuario en el contexto")
		ctx.JSON(400, gin.H{"error": "No se pudo obtener el ID de usuario"})
		return
	}
	var adminIDInt int
	switch v := adminID.(type) {
	case int:
		adminIDInt = v
	case float64:
		adminIDInt = int(v) 
	case string:
		var err error
		adminIDInt, err = strconv.Atoi(v)
		if err != nil {
			log.Println("El ID de usuario no es un número válido:", err)
			ctx.JSON(400, gin.H{"error": "El ID de usuario no es un número válido"})
			return
		}
	default:
		log.Println("Tipo de ID de usuario no válido")
		ctx.JSON(400, gin.H{"error": "Tipo de ID de usuario no válido"})
		return
	}

	fmt.Println("ID de usuario extraído:", adminIDInt)
	var request struct {
		Nombre   string  `json:"nombre"`
		Marca    string  `json:"marca"`
		Tipo     string  `json:"tipo"`
		Precio   float64 `json:"precio"`
		Densidad float64 `json:"densidad"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error decodificando el cuerpo de la solicitud: %v", err)
		ctx.JSON(400, gin.H{"error": "Cuerpo de la solicitud inválido"})
		return
	}
	log.Printf("Creando jabón: Nombre=%s, Marca=%s, Tipo=%s, Precio=%f, Densidad=%f, Usuario id=%d",
		request.Nombre, request.Marca, request.Tipo, request.Precio, request.Densidad, adminIDInt)
	if err := c.createSoap.Run(request.Nombre, request.Marca, request.Tipo, request.Precio, request.Densidad, adminIDInt); err != nil {
		log.Printf("Error creando el jabón: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Jabón creado exitosamente")
	ctx.JSON(201, gin.H{"message": "Jabón creado exitosamente"})
}
