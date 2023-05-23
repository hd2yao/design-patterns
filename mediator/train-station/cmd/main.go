package main

import (
    "design-patterns/mediator/train-station/manager"
    "design-patterns/mediator/train-station/train"
)

func main() {
    stationManager := manager.NewStationManager()
    passengerTrain := train.NewPassengerTrain(stationManager)
    freightTrain := train.NewFreightTrain(stationManager)

    passengerTrain.Arrive()
    freightTrain.Arrive()
    passengerTrain.Depart()
}
