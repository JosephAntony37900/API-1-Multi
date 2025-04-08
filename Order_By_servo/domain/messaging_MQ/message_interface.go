package messagingmq

type ServoMessagePublisher interface {
    PublishToServoQueue(codigoIdentificador string, despachoSegundos int) error
}

type OrderProcessor interface {
    ProcessOrder(codigoIdentificador string, despachoSegundos int, infrarrojoEstado string, infrarrojoTipo bool) error
    HandleInactivity(codigoIdentificador string) error
}
