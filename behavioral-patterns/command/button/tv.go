package main

import "fmt"

type TV struct {
    isRunning bool
}

func (t *TV) on() {
    t.isRunning = true
    fmt.Println("Turning tv on")
}

func (t *TV) off() {
    t.isRunning = false
    fmt.Println("Turning tv off")
}
