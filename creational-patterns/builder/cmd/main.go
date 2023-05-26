package main

import (
    "fmt"

    "design-patterns/creational-patterns/builder/concrete-builders"
    "design-patterns/creational-patterns/builder/director"
)

func main() {
    normalBuilder := concrete_builders.NewNormalBuilder()
    iglooBuilder := concrete_builders.NewIglooBuilder()

    director := director.NewDirector()

    director.SetBuilder(normalBuilder)
    normalHouse := director.BuildHouse()
    fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
    fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

    director.SetBuilder(iglooBuilder)
    iglooHouse := director.BuildHouse()
    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

}
