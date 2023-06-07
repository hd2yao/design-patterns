package strategy

import "fmt"

type WxPay struct{}

func (x *WxPay) OrderPay(px *payCtx) {
	fmt.Printf("Wx支付加工支付请求 %v\n", px.payParams)
	fmt.Println("正在使用Wx支付进行支付")
}
