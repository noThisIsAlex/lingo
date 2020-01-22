// Code generated by Lingo for table information_schema.INNODB_FOREIGN - DO NOT EDIT

package qinnodbforeign

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QInnodbForeign {
	return newQInnodbForeign(alias)
}

func New() QInnodbForeign {
	return newQInnodbForeign("")
}

func newQInnodbForeign(alias string) QInnodbForeign {
	q := QInnodbForeign{_alias: alias}
	q.id = path.NewStringPath(q, "ID")
	q.forName = path.NewStringPath(q, "FOR_NAME")
	q.refName = path.NewStringPath(q, "REF_NAME")
	q.nCols = path.NewInt64Path(q, "N_COLS")
	q.__type = path.NewInt64Path(q, "TYPE")
	return q
}

type QInnodbForeign struct {
	_alias  string
	id      path.StringPath
	forName path.StringPath
	refName path.StringPath
	nCols   path.Int64Path
	__type  path.Int64Path
}

// core.Table Functions

func (q QInnodbForeign) GetColumns() []core.Column {
	return []core.Column{
		q.id,
		q.forName,
		q.refName,
		q.nCols,
		q.__type,
	}
}

func (q QInnodbForeign) GetSQL(d core.Dialect, sql core.SQL) error {
	return path.ExpandTableWithDialect(d, q, sql)
}

func (q QInnodbForeign) GetAlias() string {
	return q._alias
}

func (q QInnodbForeign) GetName() string {
	return "INNODB_FOREIGN"
}

func (q QInnodbForeign) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbForeign) Id() path.StringPath {
	return q.id
}

func (q QInnodbForeign) ForName() path.StringPath {
	return q.forName
}

func (q QInnodbForeign) RefName() path.StringPath {
	return q.refName
}

func (q QInnodbForeign) NCols() path.Int64Path {
	return q.nCols
}

func (q QInnodbForeign) Type() path.Int64Path {
	return q.__type
}
