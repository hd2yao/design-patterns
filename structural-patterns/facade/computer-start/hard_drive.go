package main

import "fmt"

type HardDrive struct{}

func (hd *HardDrive) Read(lba int, size int) []byte {
    fmt.Println("HardDrive.Read()")
    return make([]byte, 0)
}
