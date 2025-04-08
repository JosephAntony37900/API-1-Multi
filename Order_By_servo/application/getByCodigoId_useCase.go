package application

import (
	"github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/repository"
)

type GetOrderByCodigoId struct {
	repo repository.OrderRepository
}

func NewGetOrderByCodigoId(repo repository.OrderRepository) *GetOrderByCodigoId{
	return &GetOrderByCodigoId{repo: repo}
}

func (gobci *GetOrderByCodigoId) Run(Codigo_Identificador string) (*entities.Order, error) {
	order, err := gobci.repo.FindById(Codigo_Identificador)
	if err != nil{
		return nil,err
	}
	return order,nil
}