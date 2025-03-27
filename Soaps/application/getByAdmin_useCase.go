package application

import (
	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Soaps/domain/repository"
)

type GetSoapsByAdmin struct {
	repo repository.SoapsRepository
}

func NewGetSoapsByAdmin(repo repository.SoapsRepository) *GetSoapsByAdmin {
	return &GetSoapsByAdmin{repo: repo}
}

func (gsa *GetSoapsByAdmin) Run(adminId int) ([]entities.Soaps, error) {
	soaps, err := gsa.repo.FindByAdminId(adminId)
	if err != nil {
		return nil, err
	}
	return soaps, nil
}