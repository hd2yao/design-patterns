package expression

func NewNumberExpressionFactory(val int) *numberExpression {
    return &numberExpression{val: val}
}

type numberExpression struct {
    val int
}

// Interpret 解释--返回整数值
func (numberExp *numberExpression) Interpret() int {
    return numberExp.val
}
