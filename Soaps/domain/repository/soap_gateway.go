package repository

import "github.com/JosephAntony37900/API-1-Multi/Soaps/domain/entities"

type SoapsRepository interface {
	Save(soaps entities.Soaps) error
	FindById(id int) (*entities.Soaps, error)
	GetAll() ([]entities.Soaps, error)
	Update(soaps entities.Soaps) error
	Delete (id int) error
} 