package eventbus

import (
	"GoReal/subscribers"
	"errors"
	"sync"
)

type EventBus struct {
	subscribers map[string][]chan interface{}
	mutex       sync.RWMutex
}

func New() *EventBus {
	return &EventBus{subscribers: make(map[string][]chan interface{})}
}

func (bus *EventBus) Subscribe(topic string, subscriber subscribers.Subscriber) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()

	if existingSlice, existingSliceWasFound := bus.subscribers[topic]; existingSliceWasFound {
		bus.subscribers[topic] = append(existingSlice, subscriber.GetChannel())
	} else {
		bus.subscribers[topic] = []chan interface{}{subscriber.GetChannel()}
	}
	go subscriber.Start()
}

func (bus *EventBus) Publish(topic string, information string) error {
	bus.mutex.RLock()
	defer bus.mutex.RUnlock()

	if _, existingSliceWasFound := bus.subscribers[topic]; existingSliceWasFound {
		for _, subscriber := range bus.subscribers[topic] {
			go func(subscriber chan interface{}) {
				subscriber <- information
			}(subscriber)
		}
	} else {
		return errors.New("topic does not exist")
	}

	return nil
}
