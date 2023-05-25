package product

type musket struct {
    gun
}

func NewMusketFactory() *musket {
    return &musket{}
}
