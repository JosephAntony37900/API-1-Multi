// Level_reading/infrastructure/rabbitmq/consumer.go
package rabbitmq

import (
    "github.com/JosephAntony37900/API-1-Multi/helpers"
    "github.com/streadway/amqp"
    "log"
	"fmt"
	"strconv"

	"github.com/JosephAntony37900/API-1-Multi/Level_reading/application"
	
)

// ConfigureAndConsume configura la cola y comienza a consumir mensajes.
func ConfigureAndConsume(queueName, routingKey, exchangeName string, handleMessage func(msg amqp.Delivery)) error {
    channel := helpers.GetRabbitMQChannel()
    if channel == nil {
        return logError("RabbitMQ channel is not initialized")
    }

    // Declarar el exchange
    if err := channel.ExchangeDeclare(
        exchangeName, // nombre del exchange
        "topic",      // tipo
        true,         // durable
        false,        // auto-delete
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    ); err != nil {
        return logError("Failed to declare exchange: %v", err)
    }

    // Declarar la cola
    queue, err := channel.QueueDeclare(
        queueName, // nombre de la cola
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return logError("Failed to declare queue: %v", err)
    }

    // Vincular la cola al exchange con la routing key
    if err := channel.QueueBind(
        queue.Name,   // nombre de la cola
        routingKey,   // routing key
        exchangeName, // exchange
        false,        // no-wait
        nil,          // arguments
    ); err != nil {
        return logError("Failed to bind queue: %v", err)
    }

    // Configurar el consumo
    messages, err := channel.Consume(
        queue.Name, // nombre de la cola
        "",         // consumer
        true,       // auto-ack
        false,      // exclusive
        false,      // no-local
        false,      // no-wait
        nil,        // args
    )
    if err != nil {
        return logError("Failed to register a consumer: %v", err)
    }

    // Procesar mensajes
    go func() {
        for msg := range messages {
            log.Printf("Received a message: %s", msg.Body)
            handleMessage(msg)
        }
    }()

    log.Println("Waiting for messages...")
    return nil
}

// logError simplifica el manejo de errores con logging
func logError(format string, args ...interface{}) error {
    log.Printf(format, args...)
    return fmt.Errorf(format, args...)
}

func StartLevelReadingConsumer(service *application.LevelReadingMessageService, queueName, routingKey, exchangeName string) error {
	handleMessage := func(msg amqp.Delivery) {
		log.Printf("Received a message: %s", msg.Body)

		// Parsear el mensaje recibido (nivel de lectura)
		level, err := strconv.ParseFloat(string(msg.Body), 64)
		if err != nil {
			log.Printf("Error al parsear el nivel de lectura: %v", err)
			return
		}

		// Suponiendo un ID de jabón fijo (o cambia según tus requisitos)
		idJabon := 1 // Cambia este valor si el ID del jabón se obtiene dinámicamente

		// Procesar el mensaje con el servicio de negocio
		err = service.ProcessMessage(level, idJabon)
		if err != nil {
			log.Printf("Error al procesar el mensaje: %v", err)
		}
	}

	return ConfigureAndConsume(queueName, routingKey, exchangeName, handleMessage)
}
