package strategy

import "fmt"

type ThirdPay struct{}

func (t *ThirdPay) OrderPay(px *payCtx) {
    fmt.Printf("三方支付加工支付请求 %v\n", px.payParams)
    fmt.Println("正在使用三方支付进行支付")
}
