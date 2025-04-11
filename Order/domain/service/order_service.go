package service

import (
	"fmt"
	"log"
	"time"

	"github.com/JosephAntony37900/API-1-Multi/Order/domain/entities"
	messagingmq "github.com/JosephAntony37900/API-1-Multi/Order/domain/messaging_MQ"
	"github.com/JosephAntony37900/API-1-Multi/Order/domain/repository"
)

type OrderService struct {
	repo     repository.OrderRepository
	bombaPub messagingmq.MessagePublisher
	servoPub messagingmq.MessagePublisher
}

func NewOrderService(repo repository.OrderRepository, bombaPub, servoPub messagingmq.MessagePublisher) *OrderService {
	return &OrderService{
		repo:     repo,
		bombaPub: bombaPub,
		servoPub: servoPub,
	}
}

func (service *OrderService) ProcessOrder(codigoIdentificador string, despachoSegundos int, tipo bool) error {
	log.Printf("[PROCESAR] Inicio - Codigo: %s, Tiempo: %d, Tipo: %t", codigoIdentificador, despachoSegundos, tipo)

	order, err := service.repo.FindById(codigoIdentificador)
	if err != nil && err.Error() != "no se encontró ninguna orden" {
		return fmt.Errorf("error buscando orden: %w", err)
	}

	if order == nil {
		log.Printf("[ORDEN] Creando nueva orden para %s", codigoIdentificador)
		newOrder := entities.Order{
			Codigo_Identificador: codigoIdentificador,
			Estado:               2,
			Tipo:                 tipo,
			Id_Jabon:             1,
		}
		if err := service.repo.Save(newOrder); err != nil {
			return fmt.Errorf("error creando orden: %w", err) // aquí es donde ocurre el error
		}
		order = &newOrder
	}

	var publisher messagingmq.MessagePublisher
	if tipo {
		publisher = service.bombaPub
		log.Printf("[BOMBA] Enviando a motor/bomba por %d segundos", despachoSegundos)
	} else {
		publisher = service.servoPub
		log.Printf("[SERVO] Enviando a motor/servo por %d segundos", despachoSegundos)
	}

	if err := publisher.Publish(codigoIdentificador, despachoSegundos); err != nil {
		return fmt.Errorf("error publicando el mensaje: %w", err)
	}

	go func() {
		time.Sleep(time.Duration(despachoSegundos) * time.Second)
		service.ChangeOrderState(codigoIdentificador, 1)
	}()

	return nil
}

func (service *OrderService) ChangeOrderState(codigoIdentificador string, nuevoEstado int) error {
	order, err := service.repo.FindById(codigoIdentificador)
	if err != nil {
		return fmt.Errorf("error encontrando la orden: %w", err)
	}

	if order == nil {
		log.Printf("[ORDEN] No se encontró la orden para %s. No se puede cambiar el estado.", codigoIdentificador)
		return fmt.Errorf("orden no encontrada para %s", codigoIdentificador)
	}

	order.Estado = nuevoEstado
	if err := service.repo.Update(*order); err != nil {
		return fmt.Errorf("error actualizando el estado de la orden: %w", err)
	}

	log.Printf("[ORDEN] Estado cambiado. Código: %s, Nuevo Estado: %d", codigoIdentificador, nuevoEstado)
	return nil
}

func (service *OrderService) HandleInactivity(codigoIdentificador string) error {
	order, err := service.repo.FindById(codigoIdentificador)
	if err != nil {
		log.Printf("[INACTIVIDAD] Error encontrando la orden para %s: %v", codigoIdentificador, err)
		return nil
	}

	if order == nil {
		log.Printf("[INACTIVIDAD] No se encontró ninguna orden para %s. No se cambiará estado.", codigoIdentificador)
		return nil
	}

	order.Estado = 5
	if err := service.repo.Update(*order); err != nil {
		return fmt.Errorf("error actualizando estado de inactividad: %w", err)
	}

	log.Printf("[INACTIVIDAD] Estado cambiado a 'Inactivo' para %s.", codigoIdentificador)
	return nil
}
