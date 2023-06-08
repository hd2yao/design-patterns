package main

import "fmt"

type Vehicle interface {
    Drive()
}

type Car struct{}

func (c *Car) Drive() {
    fmt.Println("Car is being driven")
}

type Driver struct {
    Age int
}
