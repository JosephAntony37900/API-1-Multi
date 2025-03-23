package application

import (
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/repository"
)

type GetByIdLevelReading struct {
	repo repository.Level_ReadingRepository
}

func NewGetByIdLevelReading(repo repository.Level_ReadingRepository) *GetByIdLevelReading{
	return &GetByIdLevelReading{repo: repo}
}

func(gbilr *GetByIdLevelReading) Run(Id int) (*entities.Level_Reading, error) {
	levelReading, err := gbilr.repo.FindById(Id)
	if err != nil {
		return nil, err
	}

	return levelReading, err
}