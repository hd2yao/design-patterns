package main

import (
    "fmt"

    "design-patterns/creational-patterns/prototype/DOM-obj/concrete-prototype"
)

func main() {
    // 职级节点--总监
    directorNode := concrete_prototype.NewElement("Director of Engineering")
    // 职级节点--研发经理
    engManagerNode := concrete_prototype.NewElement("Engineering Manager")
    engManagerNode.AddChild(concrete_prototype.NewElement("Lead Software Engineer"))
    // 研发经理是总监的下级
    directorNode.AddChild(engManagerNode)
    // 办公室经理也是总监的下级
    officeManagerNode := concrete_prototype.NewElement("Office Manager")
    directorNode.AddChild(officeManagerNode)
    fmt.Println("")
    fmt.Println("# Company Hierarchy")
    fmt.Print(directorNode)
    fmt.Println("")
    // 从研发经理节点克隆出一颗新的树
    fmt.Println("# Team Hierarchy")
    fmt.Print(engManagerNode.Clone())
}
