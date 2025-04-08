package messagingmq

type ServoMessagePublisher interface {
    PublishToServoQueue(codigoIdentificador string, despachoSegundos int) error
}

type Message struct {
    CodigoIdentificador string `json:"CodigoIdentificador"`
    Estado              string `json:"Estado"`
    Tipo                bool   `json:"Tipo"`
}

type OrderProcessor interface {
    ProcessOrder(codigoIdentificador string, despachoSegundos int, infrarrojoEstado string, infrarrojoTipo bool) error
    GetLastInfraredState(codigoIdentificador string) (*Message, error)
}
