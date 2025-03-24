package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Users/domain/repository"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"
	helpers "github.com/JosephAntony37900/API-1-Multi/helpers"
)

type LoginUser struct {
	repo repository.UserRepository
}

func NewLoginUser(repo repository.UserRepository) *LoginUser {
	return &LoginUser{repo: repo}
}

func (lu *LoginUser) Run(email string, password string) (*entities.Users, bool, error) {
	user, err := lu.repo.FindByEmail(email)
	if err != nil {
		return nil, false, fmt.Errorf("usuario no encontrado: %w", err)
	}

	fmt.Println("Contraseña guardada en la BD", user.Contraseña)

	if !helpers.ComparePassword(user.Contraseña, password) {
		return nil, false, fmt.Errorf("contraseña incorrecta")
	}

	return user, true, nil
}

