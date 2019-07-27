// Code generated by Lingo for table information_schema.SESSION_STATUS - DO NOT EDIT

package qsessionstatus

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QSessionStatus {
	return newQSessionStatus(alias)
}

func New() QSessionStatus {
	return newQSessionStatus("")
}

func newQSessionStatus(alias string) QSessionStatus {
	q := QSessionStatus{_alias: alias}
	q.variableName = path.NewStringPath(q, "VARIABLE_NAME")
	q.variableValue = path.NewStringPath(q, "VARIABLE_VALUE")
	return q
}

type QSessionStatus struct {
	_alias        string
	variableName  path.StringPath
	variableValue path.StringPath
}

// core.Table Functions

func (q QSessionStatus) GetColumns() []core.Column {
	return []core.Column{
		q.variableName,
		q.variableValue,
	}
}

func (q QSessionStatus) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QSessionStatus) GetAlias() string {
	return q._alias
}

func (q QSessionStatus) GetName() string {
	return "SESSION_STATUS"
}

func (q QSessionStatus) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QSessionStatus) VariableName() path.StringPath {
	return q.variableName
}

func (q QSessionStatus) VariableValue() path.StringPath {
	return q.variableValue
}