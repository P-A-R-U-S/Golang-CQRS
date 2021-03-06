package Test

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

const eventFake1 = "_EventFake1"
const eventFake2 = "_EventFake2"

type fakeHandler1 struct {
	name, event                                                                    string
	isOnSubscribeFired, isOnUnsubscribeFired, isExecuteFired                       bool
	isPanicFromGoroutine                                                           bool
	isPanicOnEvent, isPanicOnOnSubscribe, isPanicOnOnUnsubscribe, isPanicOnExecute bool
	isDisableMessage, isBeforeExecuteSleep, isAfterExecuteSleep                    bool
	delay                                                                          time.Duration
	argsChanges                                                                    []interface{}
}

func (h *fakeHandler1) Event() string {
	if h.isPanicOnEvent {
		if !h.isPanicFromGoroutine {
			panic(errors.New(h.event + ":Panic in Event"))
		} else {
			go func() {
				panic(errors.New(h.event + ":" + h.name + ":Panic in Event inside of goroutine."))
			}()
		}
	}

	if len(h.event) > 0 {
		return h.event
	}

	return eventFake1
}
func (h *fakeHandler1) Execute(args ...interface{}) error {
	//fmt.Printf("--> %s : %s Args before changes %d\n", h.name, h.Event(), args)

	if h.isBeforeExecuteSleep {
		time.Sleep(h.delay)
	}

	if h.isPanicOnExecute {
		if !h.isPanicFromGoroutine {
			panic(errors.New(h.event + ":Panic in Execute"))
		} else {
			go func() {
				panic(errors.New(h.event + ":" + h.name + ":Panic in Execute inside of goroutine."))
			}()
		}
	}

	for i := 0; i < len(args); i++ {
		if _, ok := args[i].(int); ok {
			args[i] = args[i].(int) + 1000
		}
	}

	h.isExecuteFired = true

	time.Sleep(time.Microsecond * 500)

	if h.isAfterExecuteSleep {
		time.Sleep(h.delay)
	}

	if !h.isDisableMessage {
		fmt.Printf("Executed: %s : %s", h.name, h.Event())
		fmt.Println()
	}

	//fmt.Printf("--> %s : %s Args after changes %d\n", h.name, h.Event(), args)
	h.argsChanges = make([]interface{}, len(args))
	for i, arg := range args {
		h.argsChanges[i] = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	}

	return nil
}
func (h *fakeHandler1) OnSubscribe() {

	if h.isPanicOnOnSubscribe {
		if !h.isPanicFromGoroutine {
			panic(errors.New(h.event + ":Panic in OnSubscribe"))
		} else {
			go func() {
				panic(errors.New(h.event + ":" + h.name + ":Panic in OnSubscribe inside of goroutine."))
			}()
		}
	}

	h.isOnSubscribeFired = true
}
func (h *fakeHandler1) OnUnsubscribe() {

	if h.isPanicOnOnUnsubscribe {
		if !h.isPanicFromGoroutine {
			panic(errors.New(h.event + ":Panic in OnUnsubscribe"))
		} else {
			go func() {
				panic(errors.New(h.event + ":" + h.name + ":Panic in OnUnsubscribe inside of goroutine."))
			}()
		}
	}

	h.isOnUnsubscribeFired = true
}

type fakeHandler2 struct {
	name, event                                                                    string
	isOnSubscribeFired, isOnUnsubscribeFired, isExecuteFired                       bool
	isPanicFromGoroutine                                                           bool
	isPanicOnEvent, isPanicOnOnSubscribe, isPanicOnOnUnsubscribe, isPanicOnExecute bool
	isDisableMessage, isBeforeExecuteSleep, isAfterExecuteSleep                    bool
	delay                                                                          time.Duration
	argsChanges                                                                    []interface{}
}

func (h *fakeHandler2) Event() string {

	if h.isPanicOnEvent {
		if !h.isPanicFromGoroutine {
			panic(errors.New(h.event + ":Panic in Event"))
		} else {
			go func() {
				panic(errors.New(h.event + ":" + h.name + ":Panic in Event inside of goroutine."))
			}()
		}
	}

	if len(h.event) > 0 {
		return h.event
	}

	return ""
}
func (h *fakeHandler2) Execute(args ...interface{}) error {
	//fmt.Printf("--> %s : %s Args before changes %d\n", h.name, h.Event(), args)

	if !h.isDisableMessage {
		fmt.Printf("Executed: %s : %s", h.name, h.Event())
		fmt.Println()
	}

	if h.isBeforeExecuteSleep {
		time.Sleep(h.delay)
	}

	if h.isPanicOnExecute {
		if !h.isPanicFromGoroutine {
			panic(errors.New(h.event + ":Panic in Execute"))
		} else {
			go func() {
				panic(errors.New(h.event + ":" + h.name + ":Panic in Execute inside of goroutine."))
			}()
		}
	}

	for i := 0; i < len(args); i++ {
		if _, ok := args[i].(int); ok {
			args[i] = args[i].(int) + 2000
		}
	}

	h.isExecuteFired = true

	time.Sleep(time.Microsecond * 500)

	if h.isAfterExecuteSleep {
		time.Sleep(h.delay)
	}

	h.argsChanges = make([]interface{}, len(args))
	for i, arg := range args {
		h.argsChanges[i] = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	}

	return nil
}
func (h *fakeHandler2) OnSubscribe() {

	if h.isPanicOnOnSubscribe {
		if !h.isPanicFromGoroutine {
			panic(errors.New(h.event + ":Panic in OnSubscribe"))
		} else {
			go func() {
				panic(errors.New(h.event + ":" + h.name + ":Panic in OnSubscribe inside of goroutine."))
			}()
		}
	}

	h.isOnSubscribeFired = true
}
func (h *fakeHandler2) OnUnsubscribe() {
	if h.isPanicOnOnUnsubscribe {
		if !h.isPanicFromGoroutine {
			panic(errors.New(h.event + ":Panic in OnUnsubscribe"))
		} else {
			go func() {
				panic(errors.New(h.event + ":" + h.name + ":Panic in OnUnsubscribe inside of goroutine."))
			}()
		}
	}

	h.isOnUnsubscribeFired = true
}
