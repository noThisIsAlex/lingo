// Code generated by Lingo for table information_schema.INNODB_SYS_TABLES - DO NOT EDIT

package qinnodbsystables

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbSysTables {
	return newQInnodbSysTables(alias)
}

func New() QInnodbSysTables {
	return newQInnodbSysTables("")
}

func newQInnodbSysTables(alias string) QInnodbSysTables {
	q := QInnodbSysTables{_alias: alias}
	q.tableId = path.NewInt64Path(q, "TABLE_ID")
	q.name = path.NewStringPath(q, "NAME")
	q.flag = path.NewIntPath(q, "FLAG")
	q.nCols = path.NewIntPath(q, "N_COLS")
	q.space = path.NewIntPath(q, "SPACE")
	q.fileFormat = path.NewStringPath(q, "FILE_FORMAT")
	q.rowFormat = path.NewStringPath(q, "ROW_FORMAT")
	q.zipPageSize = path.NewIntPath(q, "ZIP_PAGE_SIZE")
	q.spaceType = path.NewStringPath(q, "SPACE_TYPE")
	return q
}

type QInnodbSysTables struct {
	_alias      string
	tableId     path.Int64Path
	name        path.StringPath
	flag        path.IntPath
	nCols       path.IntPath
	space       path.IntPath
	fileFormat  path.StringPath
	rowFormat   path.StringPath
	zipPageSize path.IntPath
	spaceType   path.StringPath
}

// core.Table Functions

func (q QInnodbSysTables) GetColumns() []core.Column {
	return []core.Column{
		q.tableId,
		q.name,
		q.flag,
		q.nCols,
		q.space,
		q.fileFormat,
		q.rowFormat,
		q.zipPageSize,
		q.spaceType,
	}
}

func (q QInnodbSysTables) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbSysTables) GetAlias() string {
	return q._alias
}

func (q QInnodbSysTables) GetName() string {
	return "INNODB_SYS_TABLES"
}

func (q QInnodbSysTables) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbSysTables) TableId() path.Int64Path {
	return q.tableId
}

func (q QInnodbSysTables) Name() path.StringPath {
	return q.name
}

func (q QInnodbSysTables) Flag() path.IntPath {
	return q.flag
}

func (q QInnodbSysTables) NCols() path.IntPath {
	return q.nCols
}

func (q QInnodbSysTables) Space() path.IntPath {
	return q.space
}

func (q QInnodbSysTables) FileFormat() path.StringPath {
	return q.fileFormat
}

func (q QInnodbSysTables) RowFormat() path.StringPath {
	return q.rowFormat
}

func (q QInnodbSysTables) ZipPageSize() path.IntPath {
	return q.zipPageSize
}

func (q QInnodbSysTables) SpaceType() path.StringPath {
	return q.spaceType
}
