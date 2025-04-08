package application

import (
	"log"
	"fmt"

	"github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/repository"
)

type CreateOrder struct {
	repo repository.OrderRepository
}

func NewCreateOrder(repo repository.OrderRepository) *CreateOrder{
	return &CreateOrder{repo: repo}
}

func (co *CreateOrder) Run(cantidad float64, estado int, costo float64, codigoIdentificador string, tipo bool) error {
    idJabon := 1

    order, err := co.repo.FindById(codigoIdentificador)
    if err != nil && err.Error() != "no se encontró ninguna orden" {
        return fmt.Errorf("error buscando la orden: %w", err)
    }

    if order == nil {
        log.Printf("Creando nueva orden con Código_Identificador: %s", codigoIdentificador)
        newOrder := entities.Order{
            Codigo_Identificador: codigoIdentificador,
            Cantidad:             cantidad,
            Estado:               estado,
            Costo:                costo,
            Tipo:                 tipo,
            Id_Jabon:             idJabon,
        }
        return co.repo.Save(newOrder)
    }

    log.Printf("Actualizando orden existente con Código_Identificador: %s", codigoIdentificador)
    order.Cantidad = cantidad
    order.Estado = estado
    order.Costo = costo
    order.Tipo = tipo
    order.Id_Jabon = idJabon
    return co.repo.Update(*order)
}
