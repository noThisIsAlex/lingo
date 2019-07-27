// Code generated by Lingo for table information_schema.GLOBAL_STATUS - DO NOT EDIT

package qglobalstatus

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QGlobalStatus {
	return newQGlobalStatus(alias)
}

func New() QGlobalStatus {
	return newQGlobalStatus("")
}

func newQGlobalStatus(alias string) QGlobalStatus {
	q := QGlobalStatus{_alias: alias}
	q.variableName = path.NewStringPath(q, "VARIABLE_NAME")
	q.variableValue = path.NewStringPath(q, "VARIABLE_VALUE")
	return q
}

type QGlobalStatus struct {
	_alias        string
	variableName  path.StringPath
	variableValue path.StringPath
}

// core.Table Functions

func (q QGlobalStatus) GetColumns() []core.Column {
	return []core.Column{
		q.variableName,
		q.variableValue,
	}
}

func (q QGlobalStatus) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QGlobalStatus) GetAlias() string {
	return q._alias
}

func (q QGlobalStatus) GetName() string {
	return "GLOBAL_STATUS"
}

func (q QGlobalStatus) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QGlobalStatus) VariableName() path.StringPath {
	return q.variableName
}

func (q QGlobalStatus) VariableValue() path.StringPath {
	return q.variableValue
}
