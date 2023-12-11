package main

import "fmt"

type Memory struct{}

func (m *Memory) Load(position int, data []byte) {
	fmt.Println("Memory.Load()")
}
