package messagingmq

type MessagePublisher interface {
    Publish(codigoIdentificador string, despachoSegundos int) error
}

type Message struct {
    CodigoIdentificador string `json:"CodigoIdentificador"`
    Estado              string `json:"Estado"`
    Tipo                bool   `json:"Tipo"`
}

type OrderProcessor interface {
    ProcessOrder(codigoIdentificador string, despachoSegundos int, tipo bool) error
}