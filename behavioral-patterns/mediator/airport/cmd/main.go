package main

import (
    "design-patterns/behavioral-patterns/mediator/airport/airplane"
    "design-patterns/behavioral-patterns/mediator/airport/manage-tower"
)

func main() {
    tower := manage_tower.NewManageTower()
    boeing := airplane.NewBoeingPlane(tower)
    airbus := airplane.NewAirBusPlane(tower)

    boeing.Landing()
    airbus.Landing()
    boeing.Takeoff()
    airbus.Takeoff()

}
