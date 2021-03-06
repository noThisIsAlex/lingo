package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewIntWithAlias(e core.Table, name, alias string) Int {
	return Int{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt(e core.Table, name string) Int {
	return NewIntWithAlias(e, name, "")
}

type Int struct {
	entity core.Table
	name   string
	alias  string
}

func (i Int) GetParent() core.Table {
	return i.entity
}

func (i Int) GetName() string {
	return i.name
}

func (i Int) GetAlias() string {
	return i.alias
}

func (i Int) As(alias string) Int {
	i.alias = alias
	return i
}

func (i Int) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int) To(value int) core.Set {
	return expression.NewSet(i, expression.NewValue(value))
}

func (i Int) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(i, setExp)
}

func (i Int) Eq(equalTo int) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, expression.NewValue(equalTo))
}

func (i Int) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, equalTo)
}

func (i Int) NotEq(notEqualTo int) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, expression.NewValue(notEqualTo))
}

func (i Int) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int) LT(lessThan int) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, expression.NewValue(lessThan))
}

func (i Int) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int) LTOrEq(lessThanOrEqual int) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (i Int) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int) GT(greaterThan int) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (i Int) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int) GTOrEq(greaterThanOrEqual int) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (i Int) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int) IsNull() core.ComboExpression {
	return expression.NewOperator(i, operator.Null)
}

func (i Int) IsNotNull() core.ComboExpression {
	return expression.NewOperator(i, operator.NotNull)
}

func (i Int) In(values ...int) core.ComboExpression {
	return expression.NewOperator(i, operator.In, expression.NewValue(values))
}

func (i Int) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.In, values...)
}

func (i Int) NotIn(values ...int) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, expression.NewValue(values))
}

func (i Int) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, values...)
}

func (i Int) Between(first, second int) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (i Int) NotBetween(first, second int) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
