package main

import (
    "design-patterns/creational-patterns/factory/factory-method/creator"
    "fmt"
)

func main() {
    ak47, _ := creator.NewGunFactory("ak47").Create()
    musket, _ := creator.NewGunFactory("musket").Create()

    fmt.Printf("Gun: %s\n", ak47.GetName())
    fmt.Printf("Power: %d\n", ak47.GetPower())

    fmt.Printf("Gun: %s\n", musket.GetName())
    fmt.Printf("Power: %d\n", musket.GetPower())
}
