package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/repository"
)

type CreateLevelReading struct {
	repo repository.Level_ReadingRepository
}

func NewCreateLevelReading(repo repository.Level_ReadingRepository) *CreateLevelReading{
	return &CreateLevelReading{repo: repo}
}

func (clr *CreateLevelReading) Run(Fecha int, Id_Jabon int, Nivel_Jabon float64) error {
	levelReading := entities.Level_Reading{
		Fecha: Fecha,
		Id_Jabon: Id_Jabon,
		Nivel_Jabon: Nivel_Jabon,
	}

	if err := clr.repo.Save(levelReading); err != nil {
		return fmt.Errorf("error guardando el jabon: %w", err)
	}

	return nil
}