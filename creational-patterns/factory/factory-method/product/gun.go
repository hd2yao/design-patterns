package product

type gun struct {
    name  string
    power int
}

func (gun *gun) SetName(name string) {
    gun.name = name
}

func (gun *gun) SetPower(power int) {
    gun.power = power
}

func (gun *gun) GetName() string {
    return gun.name
}

func (gun *gun) GetPower() int {
    return gun.power
}
