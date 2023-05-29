package concrete_prototype

import (
	"bytes"
	"fmt"

	"design-patterns/creational-patterns/prototype/DOM-obj/interf"
)

func NewElement(text string) *element {
	return &element{
		text:     text,
		parent:   nil,
		children: make([]interf.Node, 0),
	}
}

type element struct {
	text     string
	parent   interf.Node
	children []interf.Node
}

func (e *element) String() string {
	buffer := bytes.NewBufferString(e.text)

	for _, c := range e.Children() {
		text := c.String()
		fmt.Fprintf(buffer, "\n %s", text)
	}

	return buffer.String()
}

func (e *element) Parent() interf.Node {
	return e.parent
}

func (e *element) SetParent(node interf.Node) {
	e.parent = node
}

func (e *element) Children() []interf.Node {
	return e.children
}

func (e *element) AddChild(child interf.Node) {
	copyTmp := child.Clone()
	copyTmp.SetParent(e)
	e.children = append(e.children, copyTmp)
}

func (e *element) Clone() interf.Node {
	copy := &element{
		text:     e.text,
		parent:   nil,
		children: make([]interf.Node, 0),
	}
	for _, child := range e.children {
		copy.AddChild(child)
	}
	return copy
}
