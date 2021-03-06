package operator

import (
	"github.com/weworksandbox/lingo/pkg/core"
)

type Operand int

const (
	Unknown Operand = iota

	And
	Or

	Eq
	NotEq
	Like
	NotLike

	LessThan
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual

	Null
	NotNull

	In
	NotIn

	Between
	NotBetween
)

var _names = map[Operand]string{
	And:                "AND",
	Or:                 "OR",
	Eq:                 "=",
	NotEq:              "<>",
	LessThan:           "<",
	LessThanOrEqual:    "<=",
	GreaterThan:        ">",
	GreaterThanOrEqual: ">=",
	Like:               "LIKE",
	NotLike:            "NOT LIKE",
	Null:               "IS NULL",
	NotNull:            "IS NOT NULL",
	In:                 "IN",
	NotIn:              "NOT IN",
	Between:            "BETWEEN",
	NotBetween:         "NOT BETWEEN",
}

func (o Operand) String() string {
	return _names[o]
}

func (o Operand) GetSQL(_ core.Dialect) (core.SQL, error) {
	return core.NewSQL(o.String(), nil), nil
}
