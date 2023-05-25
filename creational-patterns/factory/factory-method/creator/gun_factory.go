package creator

import (
    "fmt"

    "design-patterns/creational-patterns/factory/factory-method/interf"
    "design-patterns/creational-patterns/factory/factory-method/product"
)

func NewGunFactory(gunType string) (interf.IGun, error) {
    switch gunType {
    case "ak47":
        return product.NewAK47Factory(), nil
    case "musket":
        return product.NewMusketFactory(), nil
    default:
        return nil, fmt.Errorf("Wrong gun type passed")
    }
}
