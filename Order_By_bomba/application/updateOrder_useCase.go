package application

import (
	"fmt"

	_"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/repository"
)

type UpdateOrder struct {
	repo repository.OrderRepository
}

func NewUpdateOrder(repo repository.OrderRepository) *UpdateOrder{
	return &UpdateOrder{repo: repo}
}

func (uo *UpdateOrder) Run(Id_Jabon int, Cantidad float64, Estado int, Costo float64, Codigo_Identificador string, Tipo bool) error {
    order, err := uo.repo.FindById(Codigo_Identificador)
    if err != nil {
        return fmt.Errorf("Orden no encontrada: %w", err)
    }

    order.Id_Jabon = Id_Jabon
    order.Cantidad = Cantidad
    order.Estado = Estado
    order.Costo = Costo
    order.Tipo = Tipo
    order.Codigo_Identificador = Codigo_Identificador

    if err := uo.repo.Update(*order); err != nil {
        return fmt.Errorf("Error actualizando la orden: %w", err)
    }

    return nil
}