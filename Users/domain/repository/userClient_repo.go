package repository

import "github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"

type UserClientRepo interface {
	Save(user entities.UserClient) error
	FindByEmail(email string) (*entities.UserClient, error)

}