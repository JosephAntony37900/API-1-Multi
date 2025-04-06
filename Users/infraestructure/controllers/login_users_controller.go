package controllers

import (
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

	// Estructura para recibir el JSON de la solicitud
	var request struct {
		Email    string `json:"Email"`
		Password string `json:"Contrasena"`
	}

	// Validar el cuerpo de la solicitud
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("Error en la petición del body: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Petición del body inválida"})
		return
	}

	// Ejecutar el caso de uso para autenticar al usuario
	user, token, err := c.LoginUser.Run(request.Email, request.Password)
	if err != nil {
		log.Printf("Error en el login: %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Responder con el token y los datos del usuario
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"token":   token,
		"user": gin.H{
			"id":     user.Id,
			"nombre": user.Nombre,
			"email":  user.Email,
			"rol":    user.Id_Rol,
			"codigo_identificador": user.Codigo_Identificador,
		},
	})
}