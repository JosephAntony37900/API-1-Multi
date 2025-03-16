package application

import (
	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/repository"
)

type GetByIdSoap struct {
	repo repository.SoapsRepository
}

func NewGetByIdSoap(repo repository.SoapsRepository) *GetByIdSoap{
	return &GetByIdSoap{repo: repo}
}

func(gbis *GetByIdSoap) Run(Id int) (*entities.Soaps, error) {
	soaps, err := gbis.repo.FindById(Id)
	if err != nil {
		return nil, err
	}
	return soaps, nil
}