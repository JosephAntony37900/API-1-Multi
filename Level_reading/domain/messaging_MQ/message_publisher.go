package messagingmq

type MessagePublisher interface {
	Publish(message string, routingKey string) error
}