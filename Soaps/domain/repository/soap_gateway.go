package repository

import "github.com/JosephAntony37900/API-1-Multi/Soaps/domain/entities"

type SoapsRepository interface {
	save(soap entities.Soaps) error
	FindById(id int) (*entities.Soaps, error)
	getAll() ([]entities.Soaps, error)
	update(soap entities.Soaps) error
	delete (id int)
} 