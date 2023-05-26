package concrete_builders

import (
	"design-patterns/creational-patterns/builder/interf"
)

func NewNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (n *normalBuilder) SetWindowType() {
	n.windowType = "Wooden Window"
}

func (n *normalBuilder) SetDoorType() {
	n.doorType = "Wooden Door"
}

func (n *normalBuilder) SetNumFloor() {
	n.floor = 2
}

func (n *normalBuilder) GetHouse() interf.House {
	return interf.House{
		WindowType: n.windowType,
		DoorType:   n.doorType,
		Floor:      n.floor,
	}
}
