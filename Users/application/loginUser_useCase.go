package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/repository"
	helpers "github.com/JosephAntony37900/API-1-Multi/helpers"
)

type LoginUser struct {
	repo repository.UserRepository
}

func NewLoginUser(repo repository.UserRepository) *LoginUser {
	return &LoginUser{repo: repo}
}

func (lu *LoginUser) Run(email string, password string) (*entities.Users, string, error) {
	// Busca al usuario por su email
	user, err := lu.repo.FindByEmail(email)
	if err != nil {
		return nil, "", fmt.Errorf("usuario no encontrado: %w", err)
	}

	// Valida la contraseña ingresada
	if !helpers.ComparePassword(user.Contraseña, password) {
		return nil, "", fmt.Errorf("contraseña incorrecta")
	}

	// Genera el token JWT usando el ID del usuario
	token, err := helpers.GenerateJWT(user.Id)
	if err != nil {
		return nil, "", fmt.Errorf("error generando el token JWT: %w", err)
	}

	// Retorna la información del usuario y el token
	return user, token, nil
}