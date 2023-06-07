package strategy

type PayBehavior interface {
	OrderPay(px *payCtx)
}
