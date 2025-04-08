package rabbitmq

import (
    "encoding/json"
    "log"
 	"fmt"

    "github.com/streadway/amqp"
    "github.com/JosephAntony37900/API-1-Multi/helpers"
    "github.com/JosephAntony37900/API-1-Multi/Order_By_servo/domain/messaging_MQ"
)


type Message struct {
    CodigoIdentificador string `json:"CodigoIdentificador"` // Coincidir exactamente con ESP32
    Estado string `json:"Estado"`
    Tipo   bool   `json:"Tipo"`
}

func ConfigureAndConsume(queueName, routingKey, exchangeName string, handleMessage func(msg amqp.Delivery)) error {
	channel := helpers.GetRabbitMQChannel()
	if channel == nil {
		return fmt.Errorf("RabbitMQ channel is not initialized")
	}

	err := channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to declare exchange: %v", err)
	}

	queue, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to declare queue: %v", err)
	}

	err = channel.QueueBind(
		queue.Name,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to bind queue: %v", err)
	}

	messages, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to register a consumer: %v", err)
	}

	go func() {
		for msg := range messages {
			handleMessage(msg)
		}
	}()
	log.Println("Listening for messages...")
	return nil
}

func StartInfraredConsumer(orderProcessor messagingmq.OrderProcessor) error {
    handleMessage := func(msg amqp.Delivery) {
        log.Printf("Mensaje RAW recibido: %s", string(msg.Body))
        
        var message Message
        if err := json.Unmarshal(msg.Body, &message); err != nil {
            log.Printf("Error deserializando mensaje: %v", err)
            return
        }

        // Validación robusta
        if message.CodigoIdentificador == "" {
            log.Println("Error: Campo CodigoIdentificador vacío")
            return
        }

        log.Printf("Mensaje procesado - Codigo: %s, Estado: %s, Tipo: %t", 
            message.CodigoIdentificador, message.Estado, message.Tipo)

        // Solo procesar si hay vaso presente y es tipo polvo
        if message.Estado == "Vaso presente" && !message.Tipo {
            log.Println("Condiciones cumplidas - Activando servo")
            if err := orderProcessor.ProcessOrder(
                message.CodigoIdentificador,
                5, // Tiempo de despacho en segundos
                message.Estado,
                message.Tipo,
            ); err != nil {
                log.Printf("Error procesando orden: %v", err)
            }
        } else {
            log.Println("Condiciones NO cumplidas - No se activa servo")
        }
    }

    // Configurar el consumidor
    err := ConfigureAndConsume(
        "infrarrojo.despachador",
        "infrarrojo.topic",
        "amq.topic",
        handleMessage,
    )
    if err != nil {
        return fmt.Errorf("error configurando consumidor: %w", err)
    }

    log.Println("✅ Consumidor de infrarrojo iniciado correctamente")
    return nil
}