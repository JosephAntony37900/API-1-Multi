package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/repository"
)

type UpdateSoap struct {
	repo repository.SoapsRepository
}

func NewUpdateSoaps(repo repository.SoapsRepository) *UpdateSoap{
	return &UpdateSoap{repo: repo}
}

func(us *UpdateSoap) Run(Id int, Nombre string, Marca string, Tipo string, Precio float64, Densidad float64) error {
	soap, err := us.repo.FindById(Id)
	if err != nil {
		return fmt.Errorf("jabon no encontrado: %w", err)
	}

	soap.Densidad = Densidad
	soap.Nombre = Nombre
	soap.Marca= Marca
	soap.Precio=Precio
	soap.Tipo= Tipo

	if err := us.repo.Update(*soap); err != nil{
		return fmt.Errorf("Error actualizando el jabon: %w", err)
	}
	return nil
}
