package airplane

import (
    "design-patterns/behavioral-patterns/mediator/airport/interf"
    "fmt"
)

type boeingPlane struct {
    mediator interf.Mediator
}

func NewBoeingPlane(mediator interf.Mediator) *boeingPlane {
    return &boeingPlane{
        mediator: mediator,
    }
}

func (b *boeingPlane) Landing() {
    if !b.mediator.CanLanding(b) {
        fmt.Println("Airplane Boeing: 飞机跑道正在被占用，无法降落！")
        return
    }
    fmt.Println("Airplane Boeing: 已成功降落！")
}

func (b *boeingPlane) Takeoff() {
    fmt.Println("Airplane Boeing: 正在起飞离开跑道！")
    b.mediator.NotifyAboutDeparture()
}

func (b *boeingPlane) PermitLanding() {
    fmt.Println("Airplane Boeing: 收到指挥塔信号，允许降落，正在降落！")
    b.Landing()
}
