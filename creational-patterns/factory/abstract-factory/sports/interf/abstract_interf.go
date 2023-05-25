package interf

type ISports interface {
	MakeShoes() IShoes
	MakeShirt() IShirt
}

type IShoes interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type IShirt interface {
	SetLogo(logo string)
	SetSize(size string)
	GetLogo() string
	GetSize() string
}
