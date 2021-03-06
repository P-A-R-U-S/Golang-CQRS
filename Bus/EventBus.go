package Bus

import (
	handlers "Golang-CQRS/Handlers"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sync"
)

// EventBus implements publish/subscribe pattern: https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern
type EventBus interface {
	Publish(eventName string, args ...interface{})
	Subscribe(handler handlers.Handler) error
	Unsubscribe(eventName string) error
}

// eventBus struct
type eventBus struct {
	mtx      sync.RWMutex
	handlers map[string][]handlers.Handler
}

// Execute appropriate handlers
func (b *eventBus) Publish(eventName string, args ...interface{}) {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	if hs, ok := b.handlers[eventName]; ok {

		for _, h := range hs {

			cArgs := cloneArgs(args)

			// In case of Closure "h" variable will be reassigned before ever executed by goroutine.
			// Because if this you need to save value into variable and use this variable in closure.
			handler := h

			go func() {
				//Handle Panic in Handler.Execute.
				defer func() {
					if err := recover(); err != nil {
						log.Printf("Panic in EventBus.Publish: %s", err)
					}
				}()
				handler.Execute(cArgs...)
			}()
		}
	}
}

// Subscribe Handler
func (b *eventBus) Subscribe(h handlers.Handler) error {
	b.mtx.Lock()
	//Handle Panic on adding new handler
	defer func() {
		b.mtx.Unlock()
		if err := recover(); err != nil {
			log.Printf("Panic in EventBus.Subscribe: %s", err)
		}
	}()

	if h == nil {
		return errors.New("Handler can not be nil.")
	}

	if len(h.Event()) == 0 {
		return errors.New("Handlers with empty Event are not allowed.")
	}

	h.OnSubscribe()
	b.handlers[h.Event()] = append(b.handlers[h.Event()], h)

	return nil
}

// Unsubscribe Handler
func (b *eventBus) Unsubscribe(eventName string) error {
	b.mtx.Lock()
	//Handle Panic on adding new handler
	defer func() {
		b.mtx.Unlock()
		if err := recover(); err != nil {
			log.Printf("Panic in EventBus.Unsubscribe: %s", err)
		}
	}()

	if _, ok := b.handlers[eventName]; ok {

		for i, h := range b.handlers[eventName] {
			if h != nil {
				h.OnUnsubscribe()
				b.handlers[eventName] = append(b.handlers[eventName][:i], b.handlers[eventName][i+1:]...)
			}
		}

		return nil
	}

	return fmt.Errorf("event are not %s exist", eventName)
}

func cloneArgs(args []interface{}) []interface{} {
	cArgs := make([]interface{}, len(args))
	for i, arg := range args {
		cArgs[i] = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	}
	return cArgs
}

// New creates new EventBus
func New() EventBus {
	return &eventBus{
		handlers: make(map[string][]handlers.Handler),
	}
}
