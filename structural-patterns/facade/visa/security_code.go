package main

import "fmt"

// 复杂子系统的组成部分

func NewSecurityCode(code int) *SecurityCode {
    return &SecurityCode{
        code: code,
    }
}

type SecurityCode struct {
    code int
}

func (s *SecurityCode) checkCode(incomingCode int) error {
    if s.code != incomingCode {
        return fmt.Errorf("Security Code is incorrect")
    }
    fmt.Println("SecurityCode Verified")
    return nil
}
