package main

import "fmt"

const (
	// TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	// CounterTerroristDressType terrorist dress type
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[TerroristDressType] = newTerroristDress()
		return d.dressMap[TerroristDressType], nil
	}

	if dressType == CounterTerroristDressType {
		d.dressMap[CounterTerroristDressType] = newCounterTerroristDress()
		return d.dressMap[CounterTerroristDressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

// 获取享元工厂实例
func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
