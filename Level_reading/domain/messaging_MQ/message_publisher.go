package messagingmq

type MessagePublisher interface {
	Publish(estado string, idLectura int, codigoIdentificador string, tipo bool, routingKey string) error
}