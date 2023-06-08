package main

import "fmt"

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{
		vehicle: &Car{},
		driver:  driver,
	}
}

type CarProxy struct {
	vehicle Vehicle
	driver  *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.vehicle.Drive()
	} else {
		fmt.Println("Driver too young!")
	}
}
