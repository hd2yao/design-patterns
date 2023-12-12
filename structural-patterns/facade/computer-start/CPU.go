package main

import "fmt"

type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("CPU.Freeze()")
}

func (c *CPU) Jump(position int) {
	fmt.Println("CPU.Jump()")
}

func (c *CPU) Execute() {
	fmt.Println("CPU.Execute()")
}
