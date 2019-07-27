// Code generated by Lingo for table information_schema.INNODB_SYS_VIRTUAL - DO NOT EDIT

package qinnodbsysvirtual

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbSysVirtual {
	return newQInnodbSysVirtual(alias)
}

func New() QInnodbSysVirtual {
	return newQInnodbSysVirtual("")
}

func newQInnodbSysVirtual(alias string) QInnodbSysVirtual {
	q := QInnodbSysVirtual{_alias: alias}
	q.tableId = path.NewInt64Path(q, "TABLE_ID")
	q.pos = path.NewIntPath(q, "POS")
	q.basePos = path.NewIntPath(q, "BASE_POS")
	return q
}

type QInnodbSysVirtual struct {
	_alias  string
	tableId path.Int64Path
	pos     path.IntPath
	basePos path.IntPath
}

// core.Table Functions

func (q QInnodbSysVirtual) GetColumns() []core.Column {
	return []core.Column{
		q.tableId,
		q.pos,
		q.basePos,
	}
}

func (q QInnodbSysVirtual) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbSysVirtual) GetAlias() string {
	return q._alias
}

func (q QInnodbSysVirtual) GetName() string {
	return "INNODB_SYS_VIRTUAL"
}

func (q QInnodbSysVirtual) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbSysVirtual) TableId() path.Int64Path {
	return q.tableId
}

func (q QInnodbSysVirtual) Pos() path.IntPath {
	return q.pos
}

func (q QInnodbSysVirtual) BasePos() path.IntPath {
	return q.basePos
}