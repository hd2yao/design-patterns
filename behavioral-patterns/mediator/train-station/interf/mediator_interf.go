package interf

type Mediator interface {
    CanArrive(train Train) bool
    NotifyAboutDeparture()
}

type Train interface {
    Arrive()
    Depart()
    PermitArrival()
}
