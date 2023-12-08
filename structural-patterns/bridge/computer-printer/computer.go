package computer_printer

// 抽象

type Computer interface {
    Print()
    SetPrinter(Printer)
}
