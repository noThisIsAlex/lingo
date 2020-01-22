// Code generated by Lingo for table information_schema.INNODB_TEMP_TABLE_INFO - DO NOT EDIT

package qinnodbtemptableinfo

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QInnodbTempTableInfo {
	return newQInnodbTempTableInfo(alias)
}

func New() QInnodbTempTableInfo {
	return newQInnodbTempTableInfo("")
}

func newQInnodbTempTableInfo(alias string) QInnodbTempTableInfo {
	q := QInnodbTempTableInfo{_alias: alias}
	q.tableId = path.NewInt64Path(q, "TABLE_ID")
	q.name = path.NewStringPath(q, "NAME")
	q.nCols = path.NewIntPath(q, "N_COLS")
	q.space = path.NewIntPath(q, "SPACE")
	return q
}

type QInnodbTempTableInfo struct {
	_alias  string
	tableId path.Int64Path
	name    path.StringPath
	nCols   path.IntPath
	space   path.IntPath
}

// core.Table Functions

func (q QInnodbTempTableInfo) GetColumns() []core.Column {
	return []core.Column{
		q.tableId,
		q.name,
		q.nCols,
		q.space,
	}
}

func (q QInnodbTempTableInfo) GetSQL(d core.Dialect, sql core.SQL) error {
	return path.ExpandTableWithDialect(d, q, sql)
}

func (q QInnodbTempTableInfo) GetAlias() string {
	return q._alias
}

func (q QInnodbTempTableInfo) GetName() string {
	return "INNODB_TEMP_TABLE_INFO"
}

func (q QInnodbTempTableInfo) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbTempTableInfo) TableId() path.Int64Path {
	return q.tableId
}

func (q QInnodbTempTableInfo) Name() path.StringPath {
	return q.name
}

func (q QInnodbTempTableInfo) NCols() path.IntPath {
	return q.nCols
}

func (q QInnodbTempTableInfo) Space() path.IntPath {
	return q.space
}
