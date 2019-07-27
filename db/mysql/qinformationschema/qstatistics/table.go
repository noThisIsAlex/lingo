// Code generated by Lingo for table information_schema.STATISTICS - DO NOT EDIT

package qstatistics

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QStatistics {
	return newQStatistics(alias)
}

func New() QStatistics {
	return newQStatistics("")
}

func newQStatistics(alias string) QStatistics {
	q := QStatistics{_alias: alias}
	q.tableCatalog = path.NewStringPath(q, "TABLE_CATALOG")
	q.tableSchema = path.NewStringPath(q, "TABLE_SCHEMA")
	q.tableName = path.NewStringPath(q, "TABLE_NAME")
	q.nonUnique = path.NewInt64Path(q, "NON_UNIQUE")
	q.indexSchema = path.NewStringPath(q, "INDEX_SCHEMA")
	q.indexName = path.NewStringPath(q, "INDEX_NAME")
	q.seqInIndex = path.NewInt64Path(q, "SEQ_IN_INDEX")
	q.columnName = path.NewStringPath(q, "COLUMN_NAME")
	q.collation = path.NewStringPath(q, "COLLATION")
	q.cardinality = path.NewInt64Path(q, "CARDINALITY")
	q.subPart = path.NewInt64Path(q, "SUB_PART")
	q.packed = path.NewStringPath(q, "PACKED")
	q.nullable = path.NewStringPath(q, "NULLABLE")
	q.indexType = path.NewStringPath(q, "INDEX_TYPE")
	q.comment = path.NewStringPath(q, "COMMENT")
	q.indexComment = path.NewStringPath(q, "INDEX_COMMENT")
	return q
}

type QStatistics struct {
	_alias       string
	tableCatalog path.StringPath
	tableSchema  path.StringPath
	tableName    path.StringPath
	nonUnique    path.Int64Path
	indexSchema  path.StringPath
	indexName    path.StringPath
	seqInIndex   path.Int64Path
	columnName   path.StringPath
	collation    path.StringPath
	cardinality  path.Int64Path
	subPart      path.Int64Path
	packed       path.StringPath
	nullable     path.StringPath
	indexType    path.StringPath
	comment      path.StringPath
	indexComment path.StringPath
}

// core.Table Functions

func (q QStatistics) GetColumns() []core.Column {
	return []core.Column{
		q.tableCatalog,
		q.tableSchema,
		q.tableName,
		q.nonUnique,
		q.indexSchema,
		q.indexName,
		q.seqInIndex,
		q.columnName,
		q.collation,
		q.cardinality,
		q.subPart,
		q.packed,
		q.nullable,
		q.indexType,
		q.comment,
		q.indexComment,
	}
}

func (q QStatistics) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QStatistics) GetAlias() string {
	return q._alias
}

func (q QStatistics) GetName() string {
	return "STATISTICS"
}

func (q QStatistics) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QStatistics) TableCatalog() path.StringPath {
	return q.tableCatalog
}

func (q QStatistics) TableSchema() path.StringPath {
	return q.tableSchema
}

func (q QStatistics) TableName() path.StringPath {
	return q.tableName
}

func (q QStatistics) NonUnique() path.Int64Path {
	return q.nonUnique
}

func (q QStatistics) IndexSchema() path.StringPath {
	return q.indexSchema
}

func (q QStatistics) IndexName() path.StringPath {
	return q.indexName
}

func (q QStatistics) SeqInIndex() path.Int64Path {
	return q.seqInIndex
}

func (q QStatistics) ColumnName() path.StringPath {
	return q.columnName
}

func (q QStatistics) Collation() path.StringPath {
	return q.collation
}

func (q QStatistics) Cardinality() path.Int64Path {
	return q.cardinality
}

func (q QStatistics) SubPart() path.Int64Path {
	return q.subPart
}

func (q QStatistics) Packed() path.StringPath {
	return q.packed
}

func (q QStatistics) Nullable() path.StringPath {
	return q.nullable
}

func (q QStatistics) IndexType() path.StringPath {
	return q.indexType
}

func (q QStatistics) Comment() path.StringPath {
	return q.comment
}

func (q QStatistics) IndexComment() path.StringPath {
	return q.indexComment
}
