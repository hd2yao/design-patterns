package bank

import "design-patterns/behavioral-patterns/template/deposits/bank/interf"

func NewBankBusinessExecutor(businessHandler interf.BankBusinessHandler) *bankBusinessExecutor {
    return &bankBusinessExecutor{handler: businessHandler}
}

type bankBusinessExecutor struct {
    handler interf.BankBusinessHandler
}

// ExecuteBankBusiness 模板方法，处理银行业务
func (b *bankBusinessExecutor) ExecuteBankBusiness() {
    // 适用于与客户端单次交互的流程
    // 如果需要与客户端多次交互才能完成整个流程，每次交互的操作去调用对应模板里定义的方法就好，并不需要一个调用所用方法的模板方法
    b.handler.TakeRowNumber()
    if !b.handler.CheckVIPIdentity() {
        b.handler.WaitInHead()
    }
    b.handler.HandleBusiness()
    b.handler.Commentate()
}
