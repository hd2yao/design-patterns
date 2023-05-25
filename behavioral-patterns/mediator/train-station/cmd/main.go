package main

import (
    "design-patterns/behavioral-patterns/mediator/train-station/manager"
    train2 "design-patterns/behavioral-patterns/mediator/train-station/train"
)

func main() {
    stationManager := manager.NewStationManager()
    passengerTrain := train2.NewPassengerTrain(stationManager)
    freightTrain := train2.NewFreightTrain(stationManager)

    passengerTrain.Arrive()
    freightTrain.Arrive()
    passengerTrain.Depart()
}
