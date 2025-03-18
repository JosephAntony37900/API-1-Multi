package application

import (
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Users/domain/repository"
)

type GetUsers struct {
	repo repository.UserRepository
}

func NewGetUsers(repo repository.UserRepository) *GetUsers{
	return &GetUsers{repo: repo}
}

func (gu *GetUsers) Run() ([]entities.Users, error){
	users, err := gu.repo.FindAll()
	if err !=  nil {
		return nil, err
	}
	return users, nil
}