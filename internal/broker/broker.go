package broker

type Broker[T any] struct {
	subscribers map[chan T]any
}

func New[T any]() *Broker[T] {
	return &Broker[T]{
		subscribers: make(map[chan T]any),
	}
}

func (broker *Broker[T]) Subscribe() chan T {
	message := make(chan T)
	broker.subscribers[message] = nil
	return message
}

func (broker *Broker[T]) Unsubscribe(message chan T) {
	delete(broker.subscribers, message)
}

func (broker *Broker[T]) Publish(message T) {
	for subscriber := range broker.subscribers {
		subscriber <- message
	}
}
