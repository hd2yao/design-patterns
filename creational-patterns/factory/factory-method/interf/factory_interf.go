package interf

type Factory interface {
    Create() (IGun, error)
}

type IGun interface {
    SetName(name string)
    SetPower(power int)
    GetName() string
    GetPower() int
}
