package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/messaging_MQ"
	"github.com/JosephAntony37900/API-1-Multi/Order_By_bomba/domain/service"
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/streadway/amqp"
)


type Message struct {
    CodigoIdentificador string `json:"CodigoIdentificador"`
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

var lastInfraredState map[string]messagingmq.Message

func StartInfraredConsumer(orderProcessor messagingmq.OrderProcessor) error {
    handleMessage := func(msg amqp.Delivery) {
        var message messagingmq.Message
        if err := json.Unmarshal(msg.Body, &message); err != nil {
            log.Printf("Error deserializando mensaje: %v", err)
            return
        }

        if op, ok := orderProcessor.(*service.OrderService); ok {
            op.LastInfraredState[message.CodigoIdentificador] = message
            log.Printf("Estado infrarrojo actualizado: %+v", message)
        } else {
            log.Println("Error: OrderProcessor no es del tipo OrderService")
        }
    }

    err := ConfigureAndConsume(
        "infrarrojo.despachador",
        "infrarrojo.topic",
        "amq.topic",
        handleMessage,
    )
    if err != nil {
        return fmt.Errorf("error configurando consumidor: %w", err)
    }

    log.Println("Consumidor de infrarrojo iniciado correctamente")
    return nil
}
