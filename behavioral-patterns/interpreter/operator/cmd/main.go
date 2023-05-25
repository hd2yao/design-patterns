package main

import (
    "fmt"
    "log"

    "design-patterns/behavioral-patterns/interpreter/operator/parser"
)

func main() {
    p := parser.NewParserFactory()
    p.Parser("1 + 2 + 3 - 4")
    res := p.Result().Interpret()
    expect := 2
    if res != expect {
        log.Fatalf("error: expect %d got %d\n", expect, res)
    }
    fmt.Printf("expect: %d, got: %d\n", expect, res)
}
