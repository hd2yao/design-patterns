package strategy

func NewPayCtx(p PayBehavior) *payCtx {
	params := map[string]interface{}{
		"appId": "234fdfdngj4",
		"mchId": 123456,
	}
	return &payCtx{
		payBehavior: p,
		payParams:   params,
	}
}

type payCtx struct {
	payBehavior PayBehavior
	payParams   map[string]interface{}
}

func (px *payCtx) SetPayBehavior(p PayBehavior) {
	px.payBehavior = p
}

func (px *payCtx) Pay() {
	px.payBehavior.OrderPay(px)
}
