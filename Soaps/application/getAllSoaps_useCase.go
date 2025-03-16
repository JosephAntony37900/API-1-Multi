package application

import (
	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/repository"
)

type GetAllSoaps struct {
	repo repository.SoapsRepository
}

func NewGetAllSoaps(repo repository.SoapsRepository) *GetAllSoaps{
	return &GetAllSoaps{repo: repo}
}

func(gs *GetAllSoaps) Run() ([]entities.Soaps, error) {
	soaps, err := gs.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return soaps, nil
}