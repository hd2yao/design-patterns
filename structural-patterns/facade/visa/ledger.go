package main

import "fmt"

// 复杂子系统的组成部分

type Ledger struct{}

func (l *Ledger) makeEntry(accountID, txnType string, amount int) {
    fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
    return
}
