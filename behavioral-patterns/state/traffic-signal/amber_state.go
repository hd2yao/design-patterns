package main

import "fmt"

// 黄灯状态
type amberState struct {
	DefaultLightState
}

func NewAmberState() *amberState {
	state := &amberState{}
	state.StateName = "AMBER"
	return state
}

func (state *amberState) Light() {
	fmt.Println("黄灯亮起，请注意")
}

func (state *amberState) NextLight(light *TrafficLight) {
	light.TransitionState(NewRedState())
}
