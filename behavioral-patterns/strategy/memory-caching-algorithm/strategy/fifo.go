package strategy

import (
    "fmt"
)

type Fifo struct{}

func (f *Fifo) Evict(c *cache) {
    fmt.Println("Evicting by fifo strategy")
}
