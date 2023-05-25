package nike

import "design-patterns/creational-patterns/factory/abstract-factory/sports/product"

type nikeShoes struct {
	product.Shoe
}

func NewNikeShoes() *nikeShoes {
	shoe := product.Shoe{}
	shoe.SetLogo("Nike")
	return &nikeShoes{shoe}
}
