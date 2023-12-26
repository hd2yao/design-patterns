package main

import (
    "fmt"
    "reflect"
    "sync"
)

// Bus 总线
type Bus interface {
    Subscribe(topic string, handled interface{}) error
    Publish(topic string, args ...interface{})
}

// AsyncEventBus 异步事件总线
type AsyncEventBus struct {
    handlers map[string][]reflect.Value
    lock     sync.Mutex
}

func newAsyncEventBus() *AsyncEventBus {
    return &AsyncEventBus{
        handlers: map[string][]reflect.Value{},
        lock:     sync.Mutex{},
    }
}

func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
    bus.lock.Lock()
    defer bus.lock.Unlock()

    v := reflect.ValueOf(f)
    if v.Type().Kind() != reflect.Func {
        return fmt.Errorf("handler is not a function")
    }

    handler, ok := bus.handlers[topic]
    if !ok {
        handler = []reflect.Value{}
    }
    handler = append(handler, v)
    bus.handlers[topic] = handler

    return nil
}

func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
    handlers, ok := bus.handlers[topic]
    if !ok {
        fmt.Println("not found handlers in topic:", topic)
        return
    }

    params := make([]reflect.Value, len(args))
    for i, arg := range args {
        params[i] = reflect.ValueOf(arg)
    }

    for i := range handlers {
        go handlers[i].Call(params)
    }
}

func main() {

}
