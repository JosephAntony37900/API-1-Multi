// helpers/rabbitmq.go
package helpers

import (
    "github.com/streadway/amqp"
    "log"
)

var rabbitConn *amqp.Connection

func InitRabbitMQ(uri string) error {
    conn, err := amqp.Dial(uri)
    if err != nil {
        return err
    }
    rabbitConn = conn
    log.Println("Connected to RabbitMQ")
    return nil
}

func GetRabbitMQConnection() *amqp.Connection {
    return rabbitConn
}