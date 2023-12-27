package main

import "fmt"

type greenState struct {
    DefaultLightState
}

func NewGreenState() *greenState {
    state := &greenState{}
    state.StateName = "GREEN"
    return state
}

func (state *greenState) Light() {
    fmt.Println("绿灯亮起，请行驶")
}

func (state *greenState) NextLight(light *TrafficLight) {
    light.TransitionState(NewAmberState())
}
