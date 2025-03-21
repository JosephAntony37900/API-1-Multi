package application

import (
	"fmt"
	_"github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/repository"
)

type UpdateUser struct {
	repo repository.UserRepository
}

func NewUpdateUser(repo repository.UserRepository) *UpdateUser{
	return &UpdateUser{repo: repo}
}

func (us *UpdateUser) Run(id int, nombre string, email string, contraseña string)error{
	user, err := us.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("usuario no encontrado: %w", err)
	}

	//actualizo los campos del usuario:
	user.Nombre= nombre
	user.Email = email
	user.Contraseña= contraseña

	//guardo los cambios en el repositorio:
	if err := us.repo.Update(*user); err != nil {
		return fmt.Errorf("error actualizando el usuario: %w", err)
	}

	return nil
}