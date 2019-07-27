// Code generated by Lingo for table information_schema.INNODB_CMPMEM_RESET - DO NOT EDIT

package qinnodbcmpmemreset

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbCmpmemReset {
	return newQInnodbCmpmemReset(alias)
}

func New() QInnodbCmpmemReset {
	return newQInnodbCmpmemReset("")
}

func newQInnodbCmpmemReset(alias string) QInnodbCmpmemReset {
	q := QInnodbCmpmemReset{_alias: alias}
	q.pageSize = path.NewIntPath(q, "page_size")
	q.bufferPoolInstance = path.NewIntPath(q, "buffer_pool_instance")
	q.pagesUsed = path.NewIntPath(q, "pages_used")
	q.pagesFree = path.NewIntPath(q, "pages_free")
	q.relocationOps = path.NewInt64Path(q, "relocation_ops")
	q.relocationTime = path.NewIntPath(q, "relocation_time")
	return q
}

type QInnodbCmpmemReset struct {
	_alias             string
	pageSize           path.IntPath
	bufferPoolInstance path.IntPath
	pagesUsed          path.IntPath
	pagesFree          path.IntPath
	relocationOps      path.Int64Path
	relocationTime     path.IntPath
}

// core.Table Functions

func (q QInnodbCmpmemReset) GetColumns() []core.Column {
	return []core.Column{
		q.pageSize,
		q.bufferPoolInstance,
		q.pagesUsed,
		q.pagesFree,
		q.relocationOps,
		q.relocationTime,
	}
}

func (q QInnodbCmpmemReset) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbCmpmemReset) GetAlias() string {
	return q._alias
}

func (q QInnodbCmpmemReset) GetName() string {
	return "INNODB_CMPMEM_RESET"
}

func (q QInnodbCmpmemReset) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbCmpmemReset) PageSize() path.IntPath {
	return q.pageSize
}

func (q QInnodbCmpmemReset) BufferPoolInstance() path.IntPath {
	return q.bufferPoolInstance
}

func (q QInnodbCmpmemReset) PagesUsed() path.IntPath {
	return q.pagesUsed
}

func (q QInnodbCmpmemReset) PagesFree() path.IntPath {
	return q.pagesFree
}

func (q QInnodbCmpmemReset) RelocationOps() path.Int64Path {
	return q.relocationOps
}

func (q QInnodbCmpmemReset) RelocationTime() path.IntPath {
	return q.relocationTime
}