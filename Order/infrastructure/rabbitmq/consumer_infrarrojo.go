package rabbitmq

import (
	"fmt"
	"log"

	
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/streadway/amqp"
)

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
