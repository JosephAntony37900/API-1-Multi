package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	messagingmq "github.com/JosephAntony37900/API-1-Multi/Order/domain/messaging_MQ"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type RabbitMQServoPublisher struct{}

func NewRabbitMQServoPublisher() messagingmq.MessagePublisher {
	return &RabbitMQServoPublisher{}
}

func (p *RabbitMQServoPublisher) Publish(codigoIdentificador string, despachoSegundos int) error {
	mqttHost := os.Getenv("RABBITMQ_HOST")
	mqttUser := os.Getenv("RABBITMQ_USER")
	mqttPassword := os.Getenv("RABBITMQ_PASSWORD")

	brokerURL := fmt.Sprintf("tcp://%s:1883", mqttHost)

	opts := MQTT.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID("servo_publisher").
		SetUsername(mqttUser).
		SetPassword(mqttPassword).
		SetConnectTimeout(5 * time.Second)

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("Error conectando al broker MQTT: %w", token.Error())
	}

	message := map[string]interface{}{
		"ID":     codigoIdentificador,
		"Tiempo": despachoSegundos,
	}

	messageBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error serializando mensaje: %w", err)
	}

	topic := "motor/servo"
	qos := byte(0) // 0 = no garantiza entrega, podemos usar 1 si queremos asegurarlo

	token := client.Publish(topic, qos, false, messageBody)
	token.Wait()

	client.Disconnect(250)
	log.Printf("MENSAJE MQTT PUBLICADO AL SERVO - Topic: %s\nBody: %s", topic, string(messageBody))

	return nil
}
