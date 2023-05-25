package train

import (
    "fmt"

    "design-patterns/behavioral-patterns/mediator/train-station/interf"
)

type freightTrain struct {
    mediator interf.Mediator
}

func NewFreightTrain(mediator interf.Mediator) *freightTrain {
    return &freightTrain{mediator: mediator}
}

func (freight *freightTrain) Arrive() {
    if !freight.mediator.CanArrive(freight) {
        fmt.Println("FreightTrain: Arrival blocked, waiting")
        return
    }
    fmt.Println("FreightTrain: Arrived")
}

func (freight *freightTrain) Depart() {
    fmt.Println("FreightTrain: Leaving")
    freight.mediator.NotifyAboutDeparture()
}

func (freight *freightTrain) PermitArrival() {
    fmt.Println("FreightTrain: Arrival permitted")
    freight.Arrive()
}
