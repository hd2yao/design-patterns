package main

import "fmt"

type DefaultLightState struct {
    StateName string
}

func (state *DefaultLightState) EnterState() {
    fmt.Println("changed state to:", state.StateName)
}

func (state *DefaultLightState) CarPassingSpeed(road *TrafficLight, speed int, licensePlate string) {
    if speed > road.SpeedLimit {
        fmt.Printf("Car with license %s was speeding\n", licensePlate)
    }
}

func (tl *TrafficLight) TransitionState(newState LightState) {
    tl.State = newState
    tl.State.EnterState()
}
