package rabbitmq

import (
	"log"
	_"fmt"
    "encoding/json"
	"github.com/JosephAntony37900/API-1-Multi/helpers"
	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	channel      *amqp.Channel
	exchangeName string
}

func NewRabbitMQPublisher(exchangeName string) *RabbitMQPublisher {
	channel := helpers.GetRabbitMQChannel()
	return &RabbitMQPublisher{
		channel:      channel,
		exchangeName: exchangeName,
	}
}

func (p *RabbitMQPublisher) Publish(estado string, idLectura int, codigoIdentificador string, tipo bool, routingKey string) error {
	if p.channel == nil {
		return logError("RabbitMQ channel is not initialized")
	}

	// Construir el mensaje JSON con los atributos necesarios
	message := struct {
		Estado              string `json:"estado"`
		IdLectura           int    `json:"id_lectura"`
		CodigoIdentificador string `json:"codigo_identificador"`
		Tipo                bool   `json:"tipo"` // true = líquido, false = polvo
	}{
		Estado:              estado,
		IdLectura:           idLectura,
		CodigoIdentificador: codigoIdentificador,
		Tipo:                tipo,
	}

	// Convertir el mensaje a JSON
	messageBody, err := json.Marshal(message)
	if err != nil {
		return logError("Error al convertir el mensaje a JSON: %v", err)
	}

	// Publicar el mensaje en la cola
	err = p.channel.Publish(
		p.exchangeName, // exchange
		routingKey,     // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBody,
		},
	)
	if err != nil {
		return logError("Error al publicar mensaje: %v", err)
	}

	log.Printf("Mensaje publicado: %s", messageBody)
	return nil
}