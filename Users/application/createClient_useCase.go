package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/repository"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/services"
)

type CreateClients struct {
	repo                 repository.UserRepository
	bcrypt               services.IBcrypService
}

func NewCreateClient(repo repository.UserRepository, bcrypt services.IBcrypService) *CreateClients {
	return &CreateClients{
		repo:                 repo,
		bcrypt:               bcrypt,
	}
}

func (cu *CreateClients) Run(nombre string, email string, contraseña string, codigoIdentificador string) error {
	const Id_Rol = 1

	hashedPassword, err := cu.bcrypt.HashPassword(contraseña)
	if err != nil {
		return fmt.Errorf("error al encriptar la contraseña: %w", err)
	}

	user := entities.Users{
		Nombre:              nombre,
		Email:               email,
		Contraseña:          hashedPassword,
		Id_Rol:              Id_Rol,
		Codigo_Identificador: codigoIdentificador,
	}

	if err := cu.repo.SaveClient(user); err != nil {
		return fmt.Errorf("error al guardar el usuario cliente: %w", err)
	}

	return nil
}