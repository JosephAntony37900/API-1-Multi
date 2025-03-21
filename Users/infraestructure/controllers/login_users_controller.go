package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JosephAntony37900/API-1-Multi/Users/application"
	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	LoginUser *application.LoginUser
}

func NewLoginUserController(LoginUser *application.LoginUser) *LoginUserController {
	return &LoginUserController{LoginUser: LoginUser}
}

func (c *LoginUserController) Handle(ctx *gin.Context) {
	log.Println("Petición de login recibida")

	var request struct {
		Email    string `json:"Email"`
		Password string `json:"Contraseña"`
	}

	fmt.Println("emailsito: ", request.Email)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error en la petición del body: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Petición del body inválida"})
		return
	}

	valid, err := c.LoginUser.Run(request.Email, request.Password)
	fmt.Println("Contra: ", request.Password)
	if err != nil || !valid {
		log.Printf("Error en el login: %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login exitoso"})
}
