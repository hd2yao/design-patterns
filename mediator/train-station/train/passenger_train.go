package train

import (
    "design-patterns/mediator/train-station/interf"
    "fmt"
)

type passengerTrain struct {
    mediator interf.Mediator
}

func NewPassengerTrain(mediator interf.Mediator) *passengerTrain {
    return &passengerTrain{mediator: mediator}
}

func (passenger *passengerTrain) Arrive() {
    if !passenger.mediator.CanArrive(passenger) {
        fmt.Println("PassengerTrain: Arrival blocked, waiting")
        return
    }
    fmt.Println("PassengerTrain: Arrived")
}

func (passenger *passengerTrain) Depart() {
    fmt.Println("PassengerTrain: Leaving")
    passenger.mediator.NotifyAboutDeparture()
}

func (passenger *passengerTrain) PermitArrival() {
    fmt.Println("PassengerTrain: Arrival permitted, arriving")
    passenger.Arrive()
}
