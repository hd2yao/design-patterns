package interf

// Mediator 中介者--机场指挥塔的接口定义
type Mediator interface {
    CanLanding(airplane Airplane) bool
    NotifyAboutDeparture()
}

type Airplane interface {
    Landing()
    Takeoff()
    PermitLanding()
}
