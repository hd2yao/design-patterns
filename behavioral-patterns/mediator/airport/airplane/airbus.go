package airplane

import (
    "fmt"

    "design-patterns/behavioral-patterns/mediator/airport/interf"
)

type airBusPlane struct {
    mediator interf.Mediator
}

func NewAirBusPlane(mediator interf.Mediator) *airBusPlane {
    return &airBusPlane{
        mediator: mediator,
    }
}

func (airbus *airBusPlane) Landing() {
    if !airbus.mediator.CanLanding(airbus) {
        fmt.Println("Airplane AirBus: 飞机跑道正在被占用，无法降落！")
        return
    }
    fmt.Println("Airplane AirBus: 已成功降落！")
}

func (airbus *airBusPlane) Takeoff() {
    fmt.Println("Airplane AirBus: 正在起飞离开跑道！")
    airbus.mediator.NotifyAboutDeparture()
}

func (airbus *airBusPlane) PermitLanding() {
    fmt.Println("Airplane AirBus: 收到指挥塔信号，允许降落，正在降落！")
    airbus.Landing()
}
