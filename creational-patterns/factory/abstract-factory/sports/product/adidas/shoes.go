package adidas

import "design-patterns/creational-patterns/factory/abstract-factory/sports/product"

type adidasShoes struct {
	product.Shoe
}

func NewAdidasShoes() *adidasShoes {
	shoe := product.Shoe{}
	shoe.SetLogo("ADIDAS")
	return &adidasShoes{shoe}
}
