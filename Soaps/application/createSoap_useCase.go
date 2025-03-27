package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/repository"
)

type CreateSoap struct {
	repo repository.SoapsRepository
}

func NewCreateSoap(repo repository.SoapsRepository) *CreateSoap{
	return &CreateSoap{repo: repo}
}

func(cs *CreateSoap) Run(Nombre string, Marca string, Tipo string, Precio float64, Densidad float64, Id_Usuario_Admin  int) error {
	soap := entities.Soaps{
		Nombre: Nombre,
		Marca: Marca,
		Tipo: Tipo,
		Precio: Precio,
		Densidad: Densidad,
		Id_Usuario_Admin: Id_Usuario_Admin,
	}

	if err := cs.repo.Save(soap); err != nil {
		return fmt.Errorf("error guardando el jabon: %w", err)
	}

	return nil
}