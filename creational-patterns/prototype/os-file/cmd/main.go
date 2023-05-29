package main

import (
	"fmt"

	"design-patterns/creational-patterns/prototype/os-file/concrete-prototype"
)

func main() {
	file1 := concrete_prototype.NewFile("file1")
	file2 := concrete_prototype.NewFile("file2")
	file3 := concrete_prototype.NewFile("file3")

	folder1 := concrete_prototype.NewFolder("folder1", file1)
	folder2 := concrete_prototype.NewFolder("folder2", folder1, file2, file3)

	fmt.Println("\nPrinting hierarchy for folder2")
	folder2.Print(" ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone folder2")
	cloneFolder.Print(" ")

}
