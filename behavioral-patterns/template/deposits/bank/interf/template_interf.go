package interf

type BankBusinessHandler interface {
	TakeRowNumber()         // 排队拿号
	WaitInHead()            // 等位
	HandleBusiness()        // 处理具体业务
	Commentate()            // 对服务作出评价
	CheckVIPIdentity() bool // 钩子方法：用于在流程里判断是不是 VIP，实现类似 VIP 不用等的需求
}
