package main

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}
