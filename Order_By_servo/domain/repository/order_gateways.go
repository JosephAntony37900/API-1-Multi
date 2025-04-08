package repository

import "github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/entities"

type OrderRepository interface {
	Save(orders entities.Order) error
	Update(orders entities.Order) error
	FindById(id string) (*entities.Order, error)
}