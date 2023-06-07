package strategy

import (
    "fmt"
)

type Lfu struct{}

func (l *Lfu) Evict(c *cache) {
    fmt.Println("Evicting by lfu strategy")
}
