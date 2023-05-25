package factory

import (
	"design-patterns/creational-patterns/factory/abstract-factory/sports/interf"
	"design-patterns/creational-patterns/factory/abstract-factory/sports/product/nike"
)

type nikeFactory struct {
}

func (n *nikeFactory) MakeShoes() interf.IShoes {
	return nike.NewNikeShoes()
}

func (n *nikeFactory) MakeShirt() interf.IShirt {
	return nike.NewNikeShirt()
}
