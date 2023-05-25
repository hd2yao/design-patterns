package factory

import (
	"design-patterns/creational-patterns/factory/abstract-factory/sports/interf"
	"design-patterns/creational-patterns/factory/abstract-factory/sports/product/adidas"
)

type adidasFactory struct{}

func (a *adidasFactory) MakeShoes() interf.IShoes {
	return adidas.NewAdidasShoes()
}

func (a *adidasFactory) MakeShirt() interf.IShirt {
	return adidas.NewAdidasShirt()
}
