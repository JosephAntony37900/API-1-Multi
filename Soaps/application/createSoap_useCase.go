package application

import (
	_"fmt"

	_"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/repository"
)

type CreateSoap struct {
	repo repository.SoapsRepository
}

func NewCreateSoap(repo repository.SoapsRepository) *CreateSoap{
	return &CreateSoap{repo: repo}
}

