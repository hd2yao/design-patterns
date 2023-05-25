package main

import (
    airplane2 "design-patterns/behavioral-patterns/mediator/airport/airplane"
    "design-patterns/behavioral-patterns/mediator/airport/manage-tower"
)

func main() {
    tower := manage_tower.NewManageTower()
    boeing := airplane2.NewBoeingPlane(tower)
    airbus := airplane2.NewAirBusPlane(tower)

    boeing.Landing()
    airbus.Landing()
    boeing.Takeoff()
    airbus.Takeoff()

}
