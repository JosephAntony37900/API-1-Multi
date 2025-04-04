package rabbitmq

import (
	"encoding/json"
	"log"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/JosephAntony37900/API-1-Multi/Order/domain/service"
)

type Message struct {
	Estado string `json:"Estado"` // "No hay vaso" o "Vaso presente"
	Tipo   bool   `json:"Tipo"`   // true: líquido, false: polvo
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

func StartOrderConsumer(orderService *service.OrderService) error {
	handleMessage := func(msg amqp.Delivery) {
		var message Message
		err := json.Unmarshal(msg.Body, &message)
		if err != nil {
			log.Printf("Error deserializando el mensaje: %v", err)
			return
		}

		log.Printf("Mensaje recibido - Estado: %s, Tipo: %t", message.Estado, message.Tipo)

		err = orderService.ProcessOrder(message.Estado, message.Tipo)
		if err != nil {
			log.Printf("Error procesando la orden: %v", err)
		}
	}

	return ConfigureAndConsume("infrarrojo.despachador", "infrarrojo.topic", "amqp.topic", handleMessage)
}