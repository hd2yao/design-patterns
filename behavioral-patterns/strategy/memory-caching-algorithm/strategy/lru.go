package strategy

import (
    "fmt"
)

type Lru struct{}

func (l *Lru) Evict(c *cache) {
    fmt.Println("Eviction by lru strategy")
}
