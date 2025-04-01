package application

import (
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Order/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Order/domain/repository"
)

type CreateOrder struct {
	repo repository.OrderRepository
}

func NewCreateOrder(repo repository.OrderRepository) *CreateOrder{
	return &CreateOrder{repo: repo}
}

func (co *CreateOrder) Run(Cantidad float64 ,Estado int ,Costo float64,Codigo_Identificador string, Tipo bool) error{
    order := entities.Order {
	Cantidad: Cantidad ,
    Estado: Estado,
    Costo: Costo,
    Codigo_Identificador: Codigo_Identificador,
	Tipo: Tipo,
	}
    if err := co.repo.Save(order); err != nil{
         return fmt.Errorf("error realizando el despacho: %w", err)
	}
	return nil
}