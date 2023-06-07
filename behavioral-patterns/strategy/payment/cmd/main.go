package main

import "design-patterns/behavioral-patterns/strategy/payment/strategy"

func main() {
	wxPay := &strategy.WxPay{}
	px := strategy.NewPayCtx(wxPay)
	px.Pay()
	// 假设现在发现微信支付没钱，改用三方支付进行支付
	thPay := &strategy.ThirdPay{}
	px.SetPayBehavior(thPay)
	px.Pay()
}
