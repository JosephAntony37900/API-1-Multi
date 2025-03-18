package application

import (
	"fmt"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/repository"
	helpers "github.com/JosephAntony37900/API-1-Multi/helpers"
)

type LoginUser struct {
	repo repository.UserRepository
}

func NewLoginUser(repo repository.UserRepository) *LoginUser {
	return &LoginUser{repo: repo}
}

func (lu *LoginUser) Run(email string, password string) (bool, error) {
	user, err := lu.repo.FindByEmail(email)
	if err != nil {
		return false, fmt.Errorf("usuario no encontrado: %w", err)
	}

	if !helpers.ComparePassword(user.Contraseña, password) {
		return false, fmt.Errorf("contraseña incorrecta")
	}

	return true, nil
}
