package main

import "fmt"

// 复杂子系统的组成部分

type Notification struct{}

func (n *Notification) sendWalletCreditNotification() {
    fmt.Println("Send Wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
    fmt.Println("Sending wallet debit notification")
}
