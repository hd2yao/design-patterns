package decorator

import "design-patterns/structural-patterns/decorator/pizza/component"

type TomatoTopping struct {
	Pizza component.IPizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 7
}
