package helpers

import (
	"log"
	"github.com/streadway/amqp"
)

var rabbitConn *amqp.Connection
var rabbitChannel *amqp.Channel

func InitRabbitMQ(uri string) error {
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		return err
	}
	rabbitConn = conn
	log.Println("Connected to RabbitMQ")

	channel, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to create RabbitMQ channel: %v", err)
		return err
	}
	rabbitChannel = channel
	log.Println("RabbitMQ channel created")
	return nil
}

func GetRabbitMQChannel() *amqp.Channel {
	if rabbitChannel == nil {
		log.Println("RabbitMQ channel is not initialized")
	}
	return rabbitChannel
}

func CloseRabbitMQ() {
	if rabbitChannel != nil {
		rabbitChannel.Close()
	}
	if rabbitConn != nil {
		rabbitConn.Close()
	}
	log.Println("RabbitMQ connection closed")
}