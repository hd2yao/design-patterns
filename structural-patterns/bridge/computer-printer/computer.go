package main

// 抽象

type Computer interface {
    Print()
    SetPrinter(Printer)
}
