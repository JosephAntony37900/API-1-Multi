package messagingmq

type MessageConsumer interface {
    StartConsuming() error
}