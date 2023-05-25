package parser

import (
    expression2 "design-patterns/behavioral-patterns/interpreter/operator/expression"
    "design-patterns/behavioral-patterns/interpreter/operator/interf"
    "strconv"
    "strings"
)

func NewParserFactory() *Parser {
    return &Parser{}
}

type Parser struct {
    exp   []string
    index int
    prev  interf.Expression
}

func (p *Parser) Parser(exp string) {
    p.exp = strings.Split(exp, " ")

    for {
        if p.index >= len(p.exp) {
            return
        }
        switch p.exp[p.index] {
        case "+":
            p.prev = p.newAdditionExpression()
        case "-":
            p.prev = p.newSubtractionExpression()
        default:
            p.prev = p.newNumberExpression()
        }
    }
}

func (p *Parser) newNumberExpression() interf.Expression {
    v, _ := strconv.Atoi(p.exp[p.index])
    p.index++
    return expression2.NewNumberExpressionFactory(v)
}

func (p *Parser) newAdditionExpression() interf.Expression {
    p.index++
    return expression2.NewAdditionExpressionFactory(p.prev, p.newNumberExpression())
}

func (p *Parser) newSubtractionExpression() interf.Expression {
    p.index++
    return expression2.NewSubtractionExpressionFactory(p.prev, p.newNumberExpression())
}

func (p *Parser) Result() interf.Expression {
    return p.prev
}
