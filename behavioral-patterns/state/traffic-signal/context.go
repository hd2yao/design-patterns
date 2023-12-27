package main

type TrafficLight struct {
    State      LightState
    SpeedLimit int
}

func NewSimpleTrafficLight(speedLimit int) *TrafficLight {
    return &TrafficLight{
        SpeedLimit: speedLimit,
    }
}
