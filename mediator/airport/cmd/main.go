package main

import (
    "airport/airplane"
    "airport/manage-tower"
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
