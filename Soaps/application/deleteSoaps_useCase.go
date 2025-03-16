package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/repository"
)

type DeleteSoap struct {
	repo repository.SoapsRepository
}

func NewDeleteSoap(repo repository.SoapsRepository) *DeleteSoap{
	return &DeleteSoap{repo: repo}
}

func (ds *DeleteSoap) Run(id int) error {
	_, err := ds.repo.FindById(id)
	if err != nil {
		return fmt.Errorf("Jabon no encontrado: %w", err)
	}
	if err := ds.repo.Delete(id); err != nil{
		return fmt.Errorf("Error eliminando el jabon: %w", err)
	}

	return nil
}