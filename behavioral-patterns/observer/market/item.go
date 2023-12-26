package main

import "fmt"

func newItem(name string) *Item {
    return &Item{name: name}
}

type Item struct {
    observerList []Observer
    name         string
    inStock      bool
}

func (i *Item) updateAvailability() {
    fmt.Printf("Item %s is now in stock\n", i.name)
    i.inStock = true
    i.notifyAll()
}

func (i *Item) register(observer Observer) {
    i.observerList = append(i.observerList, observer)
}

func (i *Item) deregister(observer Observer) {
    i.observerList = removeFromSlice(i.observerList, observer)
}

func (i *Item) notifyAll() {
    for _, observer := range i.observerList {
        observer.update(i.name)
    }
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
    observerListLength := len(observerList)
    for i, observer := range observerList {
        if observerToRemove.getID() == observer.getID() {
            observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
            return observerList[:observerListLength-1]
        }
    }
    return observerList
}
