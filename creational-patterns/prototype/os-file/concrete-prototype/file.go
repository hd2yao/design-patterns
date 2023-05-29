package concrete_prototype

import (
	"fmt"

	"design-patterns/creational-patterns/prototype/os-file/interf"
)

func NewFile(name string) *file {
	return &file{name: name}
}

type file struct {
	name string
}

func (f *file) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) Clone() interf.Inode {
	return &file{name: f.name + "_clone"}
}
