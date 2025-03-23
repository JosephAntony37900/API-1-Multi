package application

import (
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/repository"
)

type GetLevelReading struct {
	repo repository.Level_ReadingRepository
}

func NewGetLevelReading(repo repository.Level_ReadingRepository) *GetLevelReading{
	return &GetLevelReading{repo: repo}
}

func(glr *GetLevelReading) Run() ([]entities.Level_Reading, error) {
    levelReading, err := glr.repo.GetAll()
    if err != nil {
		return nil, err
	}
	return levelReading, nil
}