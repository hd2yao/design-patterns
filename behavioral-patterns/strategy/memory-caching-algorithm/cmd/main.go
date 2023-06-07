package main

import (
    "design-patterns/behavioral-patterns/strategy/memory-caching-algorithm/strategy"
)

func main() {
    lfu := &strategy.Lfu{}
    cache := strategy.NewCache(lfu)

    cache.Add("a", "1")
    cache.Add("b", "2")

    cache.Add("c", "3")

    lru := &strategy.Lru{}
    cache.SetEvictionAlgo(lru)

    cache.Add("d", "4")

    fifo := &strategy.Fifo{}
    cache.SetEvictionAlgo(fifo)

    cache.Add("e", "5")
}
