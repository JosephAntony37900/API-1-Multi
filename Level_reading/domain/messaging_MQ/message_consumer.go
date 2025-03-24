package messagingmq

type MessageConsumer interface {
    StartConsuming() error
    ProcessMessage(level float64) error // Nuevo método para manejar el procesamiento
}