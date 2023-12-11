package main

import "fmt"

// 复杂子系统的组成部分

func NewAccount(accountName string) *Account {
    return &Account{
        name: accountName,
    }
}

type Account struct {
    name string
}

func (a *Account) checkAccount(accountName string) error {
    if a.name != accountName {
        return fmt.Errorf("Account Name is incorrect")
    }
    fmt.Println("Account Verified")
    return nil
}
