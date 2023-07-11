package main

import (
    "fmt"
    "strings"
)

type CompositeOrganization struct {
    orgName string
    depth   int
    list    []Organization
}

func NewCompositeOrganization(name string, depth int) *CompositeOrganization {
    return &CompositeOrganization{name, depth, []Organization{}}
}

func (c *CompositeOrganization) add(org Organization) {
    if c == nil {
        return
    }
    c.list = append(c.list, org)
}

func (c *CompositeOrganization) remove(org Organization) {
    if c == nil {
        return
    }
    for i, val := range c.list {
        if val == org {
            c.list = append(c.list[:i], c.list[i+1:]...)
            return
        }
    }
    return
}

func (c *CompositeOrganization) display() {
    if c == nil {
        return
    }
    fmt.Println(strings.Repeat("-", c.depth*2), " ", c.orgName)
    for _, val := range c.list {
        val.display()
    }
}

func (c *CompositeOrganization) duty() {
    if c == nil {
        return
    }

    for _, val := range c.list {
        val.duty()
    }
}
