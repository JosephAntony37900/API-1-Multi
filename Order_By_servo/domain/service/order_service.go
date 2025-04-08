package service

import (
	"fmt"
	"log"
	"time"

	"github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/entities"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/messaging_MQ"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/repository"
)

type OrderService struct {
    repo       repository.OrderRepository
    publisher  messagingmq.ServoMessagePublisher
}

func NewOrderService(repo repository.OrderRepository, publisher messagingmq.ServoMessagePublisher) *OrderService {
    return &OrderService{
        repo:      repo,
        publisher: publisher,
    }
}

func (service *OrderService) ProcessOrder(codigoIdentificador string, despachoSegundos int, infrarrojoEstado string, infrarrojoTipo bool) error {
    log.Printf("[PROCESAR] Inicio - Codigo: %s, Estado: %s, Tipo: %t", 
        codigoIdentificador, infrarrojoEstado, infrarrojoTipo)

    // Condición corregida y más clara
    vasoPresente := infrarrojoEstado == "Vaso presente"
    esPolvo := !infrarrojoTipo
    
    if vasoPresente && esPolvo {
        log.Println("[PROCESAR] Condiciones CUMPLIDAS - Vaso presente y tipo polvo")
        
        order, err := service.repo.FindById(codigoIdentificador)
        if err != nil && err.Error() != "no se encontró ninguna orden" {
            return fmt.Errorf("error buscando orden: %w", err)
        }

        if order == nil {
            log.Printf("[ORDEN] Creando nueva orden para %s", codigoIdentificador)
            newOrder := entities.Order{
                Codigo_Identificador: codigoIdentificador,
                Estado:               2, // En proceso
                Tipo:                 false,
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

        // Programar cierre después del tiempo
        go func() {
            time.Sleep(time.Duration(despachoSegundos) * time.Second)
            log.Printf("[ORDEN] Despacho completado para %s", codigoIdentificador)
            service.ChangeOrderState(codigoIdentificador, 1) // Resuelta
        }()

        return nil
    }

    log.Printf("[PROCESAR] Condiciones NO CUMPLIDAS - Vaso presente: %t, Es polvo: %t", 
        vasoPresente, esPolvo)
    return nil
}



func (service *OrderService) ChangeOrderState(codigoIdentificador string, nuevoEstado int) error {
    order, err := service.repo.FindById(codigoIdentificador)
    if err != nil {
        return fmt.Errorf("error encontrando la orden: %w", err)
    }

    order.Estado = nuevoEstado
    if err := service.repo.Update(*order); err != nil {
        return fmt.Errorf("error actualizando el estado de la orden: %w", err)
    }

    log.Printf("Estado de la orden con Codigo_Identificador %s cambiado a %d.", codigoIdentificador, nuevoEstado)
    return nil
}

func (service *OrderService) HandleInactivity(codigoIdentificador string) error {
    order, err := service.repo.FindById(codigoIdentificador)
    if err != nil {
        log.Printf("Error encontrando la orden con Código_Identificador %s: %v", codigoIdentificador, err)
        return nil
    }

    if order == nil {
        log.Printf("No se encontró ninguna orden con el Código_Identificador %s. No es necesario cambiar estado a Inactivo.", codigoIdentificador)
        return nil
    }

    order.Estado = 5 
    if err := service.repo.Update(*order); err != nil {
        return fmt.Errorf("error actualizando el estado de la orden: %w", err)
    }

    log.Printf("Estado de la orden con Codigo_Identificador %s cambiado a Inactivo.", codigoIdentificador)
    return nil
}
