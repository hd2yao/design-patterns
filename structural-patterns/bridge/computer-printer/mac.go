package computer_printer

import "fmt"

// 精确抽象

type Mac struct {
    printer Printer
}

func (m *Mac) Print() {
    fmt.Println("Print request for mac")
    m.printer.PrintFile()
}

func (m *Mac) SetPrinter(printer Printer) {
    m.printer = printer
}
