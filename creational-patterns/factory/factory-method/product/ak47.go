package product

type ak47 struct {
    gun
}

func NewAK47Factory() *ak47 {
    return &ak47{}
}
