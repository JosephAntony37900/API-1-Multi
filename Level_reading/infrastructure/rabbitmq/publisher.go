package rabbitmq

import (
	"log"
	_"fmt"

	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	channel *amqp.Channel
	exchangeName string
}

func NewRabbitMQPublisher(exchangeName string) *RabbitMQPublisher {
	channel := helpers.GetRabbitMQChannel()
	return &RabbitMQPublisher{
		channel: channel,
		exchangeName: exchangeName,
	}
}

func (p *RabbitMQPublisher) Publish(message string, routingKey string) error {
	if p.channel == nil {
		return logError("RabbitMQ channel is not initialized")
	}

	err := p.channel.Publish(
		p.exchangeName, // exchange
		routingKey,     // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return logError("Error al publicar mensaje: %v", err)
	}

	log.Printf("Mensaje publicado: %s", message)
	return nil
}