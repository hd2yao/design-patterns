package expression

import (
    "design-patterns/behavioral-patterns/interpreter/operator/interf"
)

func NewAdditionExpressionFactory(left, right interf.Expression) *additionExpression {
    return &additionExpression{
        left:  left,
        right: right,
    }
}

type additionExpression struct {
    left, right interf.Expression
}

func (additionExp *additionExpression) Interpret() int {
    return additionExp.left.Interpret() + additionExp.right.Interpret()
}
