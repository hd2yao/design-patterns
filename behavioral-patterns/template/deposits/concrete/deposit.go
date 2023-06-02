package concrete

import "fmt"

type DepositBusinessHandler struct {
	*DefaultBusinessHandler
	userVip bool
}

func (*DepositBusinessHandler) HandleBusiness() {
	fmt.Println("账户存储很多万人民币...")
}
