package interf

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

type House struct {
	WindowType string
	DoorType   string
	Floor      int
}
