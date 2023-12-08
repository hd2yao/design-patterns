package computer_printer

import "fmt"

// 精确抽象

type Windows struct {
    printer Printer
}

func (w *Windows) Print() {
    fmt.Println("Print request for windows")
    w.printer.PrintFile()
}

func (w *Windows) SetPrinter(printer Printer) {
    w.printer = printer
}
