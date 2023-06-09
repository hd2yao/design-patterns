package main

import (
	"fmt"

	"design-patterns/structural-patterns/decorator/pizza/component"
	"design-patterns/structural-patterns/decorator/pizza/decorator"
)

func main() {
	pizza := &component.VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
