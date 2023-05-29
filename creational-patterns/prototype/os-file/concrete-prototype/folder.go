package concrete_prototype

import (
	"fmt"

	"design-patterns/creational-patterns/prototype/os-file/interf"
)

func NewFolder(name string, children ...interf.Inode) *folder {
	return &folder{name: name, children: children}
}

type folder struct {
	children []interf.Inode
	name     string
}

func (f *folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.Print(indentation + indentation)
	}
}

func (f *folder) Clone() interf.Inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tmpChildren []interf.Inode
	for _, child := range f.children {
		cloneChild := child.Clone()
		tmpChildren = append(tmpChildren, cloneChild)
	}
	cloneFolder.children = tmpChildren
	return cloneFolder
}
