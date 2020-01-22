// Code generated by Lingo for table information_schema.INNODB_FT_BEING_DELETED - DO NOT EDIT

package qinnodbftbeingdeleted

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QInnodbFtBeingDeleted {
	return newQInnodbFtBeingDeleted(alias)
}

func New() QInnodbFtBeingDeleted {
	return newQInnodbFtBeingDeleted("")
}

func newQInnodbFtBeingDeleted(alias string) QInnodbFtBeingDeleted {
	q := QInnodbFtBeingDeleted{_alias: alias}
	q.docId = path.NewInt64Path(q, "DOC_ID")
	return q
}

type QInnodbFtBeingDeleted struct {
	_alias string
	docId  path.Int64Path
}

// core.Table Functions

func (q QInnodbFtBeingDeleted) GetColumns() []core.Column {
	return []core.Column{
		q.docId,
	}
}

func (q QInnodbFtBeingDeleted) GetSQL(d core.Dialect, sql core.SQL) error {
	return path.ExpandTableWithDialect(d, q, sql)
}

func (q QInnodbFtBeingDeleted) GetAlias() string {
	return q._alias
}

func (q QInnodbFtBeingDeleted) GetName() string {
	return "INNODB_FT_BEING_DELETED"
}

func (q QInnodbFtBeingDeleted) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbFtBeingDeleted) DocId() path.Int64Path {
	return q.docId
}
