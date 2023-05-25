package adidas

import "design-patterns/creational-patterns/factory/abstract-factory/sports/product"

type adidasShirt struct {
	product.Shirt
}

func NewAdidasShirt() *adidasShirt {
	shirt := product.Shirt{}
	shirt.SetLogo("ADIDAS")
	return &adidasShirt{shirt}
}
