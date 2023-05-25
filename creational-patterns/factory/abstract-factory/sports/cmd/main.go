package main

import (
	"fmt"

	"design-patterns/creational-patterns/factory/abstract-factory/sports/factory"
	"design-patterns/creational-patterns/factory/abstract-factory/sports/interf"
)

func main() {
	adidasFactory, _ := factory.GetSportsFactory("adidas")
	nikeFactory, _ := factory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoes()
	nikeShoe.SetSize(43)
	nikeShirt := nikeFactory.MakeShirt()
	nikeShirt.SetSize("L")

	adidasShoe := adidasFactory.MakeShoes()
	adidasShoe.SetSize(41)
	adidasShirt := adidasFactory.MakeShirt()
	adidasShirt.SetSize("XL")

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s interf.IShoes) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}

func printShirtDetails(s interf.IShirt) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %s\n", s.GetSize())
}
