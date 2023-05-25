package expression

import (
    "design-patterns/behavioral-patterns/interpreter/operator/interf"
)

func NewSubtractionExpressionFactory(left, right interf.Expression) *subtractionExpression {
    return &subtractionExpression{
        left:  left,
        right: right,
    }
}

type subtractionExpression struct {
    left, right interf.Expression
}

func (subtractionExp *subtractionExpression) Interpret() int {
    return subtractionExp.left.Interpret() - subtractionExp.right.Interpret()
}
