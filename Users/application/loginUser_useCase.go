package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/repository"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/services"
)

type LoginUser struct {
	repo   repository.UserRepository
	jwt    services.TokenManager
	bcrypt services.IBcrypService
}

func NewLoginUser(repo repository.UserRepository, jwt services.TokenManager, bcrypt services.IBcrypService) *LoginUser {
	return &LoginUser{
		repo:   repo,
		jwt:    jwt,
		bcrypt: bcrypt,
	}
}

func (lu *LoginUser) Run(email string, password string) (*entities.Users, string, error) {
	user, err := lu.repo.FindByEmail(email)
	if err != nil {
		return nil, "", fmt.Errorf("credenciales inválidas")
	}

	if !lu.bcrypt.ComparePasswords(user.Contraseña, password) {
		return nil, "", fmt.Errorf("credenciales inválidas")
	}

	token, err := lu.jwt.GenerateToken(user.Id)
	if err != nil {
		return nil, "", fmt.Errorf("error generando token: %w", err)
	}

	return user, token, nil
}