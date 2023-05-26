package director

import "design-patterns/creational-patterns/builder/interf"

func NewDirector() *director {
	return &director{}
}

type director struct {
	builder interf.IBuilder
}

func (d *director) SetBuilder(b interf.IBuilder) {
	d.builder = b
}

func (d *director) BuildHouse() interf.House {
	d.builder.SetWindowType()
	d.builder.SetDoorType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
