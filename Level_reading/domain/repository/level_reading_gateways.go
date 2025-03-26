package repository

import "github.com/JosephAntony37900/API-1-Multi/Level_reading/domain/entities"

type Level_ReadingRepository interface {
	Save(level_Reading entities.Level_Reading) error
	FindById(id int) (*entities.Level_Reading, error)
	GetAll() ([]entities.Level_Reading, error)
	GetLast() (*entities.Level_Reading, error)
	SaveWithReturnId(level_Reading entities.Level_Reading) (int, error)
	
}