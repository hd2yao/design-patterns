package factory

import (
	"fmt"

	"design-patterns/creational-patterns/factory/abstract-factory/sports/interf"
)

func GetSportsFactory(brand string) (interf.ISports, error) {
	if brand == "adidas" {
		return &adidasFactory{}, nil
	}

	if brand == "nike" {
		return &nikeFactory{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
