package rabbitmq

import (
    "github.com/JosephAntony37900/API-1-Multi/helpers"
    "github.com/streadway/amqp"
    "log"
	"fmt"
	_"strconv"
    "encoding/json"

	"github.com/JosephAntony37900/API-1-Multi/Level_reading/application"
	
)

type Message struct {
    ID    string  `json:"ID"`
    Nivel float64 `json:"Nivel"`
    Tipo  bool    `json:"Tipo"`
}

func ConfigureAndConsume(queueName, routingKey, exchangeName string, handleMessage func(msg amqp.Delivery)) error {
    channel := helpers.GetRabbitMQChannel()
    if channel == nil {
        return logError("RabbitMQ channel is not initialized")
    }

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

    if err := channel.QueueBind(
        queue.Name,   // nombre de la cola
        routingKey,   // routing key
        exchangeName, // exchange
        false,        // no-wait
        nil,          // arguments
    ); err != nil {
        return logError("Failed to bind queue: %v", err)
    }

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

    go func() {
        for msg := range messages {
            log.Printf("Received a message: %s", msg.Body)
            handleMessage(msg)
        }
    }()

    log.Println("Waiting for messages...")
    return nil
}

func logError(format string, args ...interface{}) error {
    log.Printf(format, args...)
    return fmt.Errorf(format, args...)
}

func StartLevelReadingConsumer(service *application.LevelReadingMessageService, queueName, routingKey, exchangeName string) error {
    handleMessage := func(msg amqp.Delivery) {
        log.Printf("Received a message: %s", msg.Body)

        var message Message
        err := json.Unmarshal(msg.Body, &message)
        if err != nil {
            log.Printf("Error al deserializar el mensaje: %v", err)
            return
        }

        log.Printf("Código identificador: %s", message.ID)
        log.Printf("Nivel de lectura: %.2f%%", message.Nivel)
        log.Printf("Tipo: %t (true = líquido, false = polvo)", message.Tipo)

        idJabon := 1 
        err = service.ProcessMessage(message.Nivel, idJabon, message.ID, message.Tipo)
        if err != nil {
            log.Printf("Error al procesar el mensaje: %v", err)
        }
    }

    return ConfigureAndConsume(queueName, routingKey, exchangeName, handleMessage)
}