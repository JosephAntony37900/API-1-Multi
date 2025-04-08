package service

import (
	"fmt"
	"log"
	"time"

	"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/messaging_MQ"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/repository"
)

type OrderService struct {
    repo       repository.OrderRepository
    publisher  messagingmq.ServoMessagePublisher
    LastInfraredState map[string]messagingmq.Message
}

func NewOrderService(repo repository.OrderRepository, publisher messagingmq.ServoMessagePublisher) *OrderService {
    return &OrderService{
        repo:      repo,
        publisher: publisher,
        LastInfraredState: make(map[string]messagingmq.Message),
    }
}

func (service *OrderService) GetLastInfraredState(codigoIdentificador string) (*messagingmq.Message, error) {
    message, exists := service.LastInfraredState[codigoIdentificador]
    if !exists {
        return nil, fmt.Errorf("estado infrarrojo no encontrado para el Código %s", codigoIdentificador)
    }
    return &message, nil
}


func (service *OrderService) ProcessOrder(codigoIdentificador string, despachoSegundos int, _ string, tipo bool) error {
    log.Printf("[PROCESAR] Inicio - Codigo: %s, Tiempo: %d, Tipo: %t", codigoIdentificador, despachoSegundos, tipo)

    infraredMessage, exists := service.LastInfraredState[codigoIdentificador]
    if !exists {
        log.Printf("[PROCESAR] Estado infrarrojo no encontrado para el Código %s. No se puede procesar la orden.", codigoIdentificador)
        return fmt.Errorf("estado infrarrojo no encontrado para el Código %s", codigoIdentificador)
    }

    vasoPresente := infraredMessage.Estado == "Vaso presente"
    esLiquido := infraredMessage.Tipo 

    if vasoPresente && esLiquido {
        log.Println("[PROCESAR] Condiciones CUMPLIDAS - Vaso presente y tipo líquido")

        order, err := service.repo.FindById(codigoIdentificador)
        if err != nil && err.Error() != "no se encontró ninguna orden" {
            return fmt.Errorf("error buscando orden: %w", err)
        }

        if order == nil {
            log.Printf("[ORDEN] Creando nueva orden para %s", codigoIdentificador)
            newOrder := entities.Order{
                Codigo_Identificador: codigoIdentificador,
                Estado:               2, 
                Tipo:                 true, 
            }
            if err := service.repo.Save(newOrder); err != nil {
                return fmt.Errorf("error creando orden: %w", err)
            }
            order = &newOrder
        }

        log.Printf("[SERVO] Enviando comando de apertura por %d segundos", despachoSegundos)
        if err := service.publisher.PublishToServoQueue(codigoIdentificador, despachoSegundos); err != nil {
            log.Printf("[ERROR] Publicando al servo: %v", err)
            return err
        }

        go func() {
            time.Sleep(time.Duration(despachoSegundos) * time.Second)
            log.Printf("[ORDEN] Despacho completado para %s", codigoIdentificador)
            service.ChangeOrderState(codigoIdentificador, 1) 
        }()
        return nil
    }

    log.Printf("[PROCESAR] Condiciones NO CUMPLIDAS - Vaso presente: %t, Es líquido: %t", vasoPresente, esLiquido)
    return fmt.Errorf("condiciones no cumplidas para el Código %s", codigoIdentificador)
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
