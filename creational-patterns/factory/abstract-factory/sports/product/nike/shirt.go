package nike

import "design-patterns/creational-patterns/factory/abstract-factory/sports/product"

type nikeShirt struct {
	product.Shirt
}

func NewNikeShirt() *nikeShirt {
	shirt := product.Shirt{}
	shirt.SetLogo("NIKE")
	return &nikeShirt{shirt}
}
