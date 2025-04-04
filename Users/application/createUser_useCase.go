package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/repository"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/services"
)

type CreateUsers struct {
	repo   repository.UserRepository
	bcrypt services.IBcrypService
}

func NewCreateUser(repo repository.UserRepository, bcrypt services.IBcrypService) *CreateUsers {
	return &CreateUsers{
		repo:   repo,
		bcrypt: bcrypt,
	}
}

func (cu *CreateUsers) Run(nombre string, email string, contraseña string, Codigo_Identificador string) error {
	hashedPassword, err := cu.bcrypt.HashPassword(contraseña)
	if err != nil {
		return fmt.Errorf("error al encriptar la contraseña: %w", err)
	}

	user := entities.Users{
		Nombre:              nombre,
		Email:               email,
		Contraseña:          hashedPassword,
		Codigo_Identificador: Codigo_Identificador,
	}

	if err := cu.repo.Save(user); err != nil {
		return fmt.Errorf("error al guardar el usuario: %w", err)
	}
	return nil
}