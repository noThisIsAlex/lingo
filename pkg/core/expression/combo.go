package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

type ComboExpression struct {
	exp core.Expression
}

func NewComboExpression(exp core.Expression) *ComboExpression {
	return &ComboExpression{
		exp: exp,
	}
}

func (c ComboExpression) And(exp core.Expression) core.ComboExpression {
	return NewOperator(c.exp, operator.And, exp)
}

func (c ComboExpression) Or(exp core.Expression) core.ComboExpression {
	return NewOperator(c.exp, operator.Or, exp)
}
