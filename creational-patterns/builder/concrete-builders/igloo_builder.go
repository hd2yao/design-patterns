package concrete_builders

import "design-patterns/creational-patterns/builder/interf"

func NewIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (i *iglooBuilder) SetWindowType() {
	i.windowType = "Snow Window"
}

func (i *iglooBuilder) SetDoorType() {
	i.doorType = "Snow Door"
}

func (i *iglooBuilder) SetNumFloor() {
	i.floor = 1
}

func (i *iglooBuilder) GetHouse() interf.House {
	return interf.House{
		DoorType:   i.doorType,
		WindowType: i.windowType,
		Floor:      i.floor,
	}
}
