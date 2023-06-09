package decorator

import "design-patterns/structural-patterns/decorator/pizza/component"

type CheeseTopping struct {
	Pizza component.IPizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}
