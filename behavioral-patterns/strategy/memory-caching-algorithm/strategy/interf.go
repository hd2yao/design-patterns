package strategy

type EvictionAlgo interface {
    Evict(c *cache)
}
