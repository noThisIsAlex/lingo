// Code generated by Lingo for table information_schema.PARTITIONS - DO NOT EDIT

package qpartitions

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QPartitions {
	return newQPartitions(alias)
}

func New() QPartitions {
	return newQPartitions("")
}

func newQPartitions(alias string) QPartitions {
	q := QPartitions{_alias: alias}
	q.tableCatalog = path.NewStringPath(q, "TABLE_CATALOG")
	q.tableSchema = path.NewStringPath(q, "TABLE_SCHEMA")
	q.tableName = path.NewStringPath(q, "TABLE_NAME")
	q.partitionName = path.NewStringPath(q, "PARTITION_NAME")
	q.subpartitionName = path.NewStringPath(q, "SUBPARTITION_NAME")
	q.partitionOrdinalPosition = path.NewInt64Path(q, "PARTITION_ORDINAL_POSITION")
	q.subpartitionOrdinalPosition = path.NewInt64Path(q, "SUBPARTITION_ORDINAL_POSITION")
	q.partitionMethod = path.NewStringPath(q, "PARTITION_METHOD")
	q.subpartitionMethod = path.NewStringPath(q, "SUBPARTITION_METHOD")
	q.partitionExpression = path.NewStringPath(q, "PARTITION_EXPRESSION")
	q.subpartitionExpression = path.NewStringPath(q, "SUBPARTITION_EXPRESSION")
	q.partitionDescription = path.NewStringPath(q, "PARTITION_DESCRIPTION")
	q.tableRows = path.NewInt64Path(q, "TABLE_ROWS")
	q.avgRowLength = path.NewInt64Path(q, "AVG_ROW_LENGTH")
	q.dataLength = path.NewInt64Path(q, "DATA_LENGTH")
	q.maxDataLength = path.NewInt64Path(q, "MAX_DATA_LENGTH")
	q.indexLength = path.NewInt64Path(q, "INDEX_LENGTH")
	q.dataFree = path.NewInt64Path(q, "DATA_FREE")
	q.createTime = path.NewTimePath(q, "CREATE_TIME")
	q.updateTime = path.NewTimePath(q, "UPDATE_TIME")
	q.checkTime = path.NewTimePath(q, "CHECK_TIME")
	q.checksum = path.NewInt64Path(q, "CHECKSUM")
	q.partitionComment = path.NewStringPath(q, "PARTITION_COMMENT")
	q.nodegroup = path.NewStringPath(q, "NODEGROUP")
	q.tablespaceName = path.NewStringPath(q, "TABLESPACE_NAME")
	return q
}

type QPartitions struct {
	_alias                      string
	tableCatalog                path.StringPath
	tableSchema                 path.StringPath
	tableName                   path.StringPath
	partitionName               path.StringPath
	subpartitionName            path.StringPath
	partitionOrdinalPosition    path.Int64Path
	subpartitionOrdinalPosition path.Int64Path
	partitionMethod             path.StringPath
	subpartitionMethod          path.StringPath
	partitionExpression         path.StringPath
	subpartitionExpression      path.StringPath
	partitionDescription        path.StringPath
	tableRows                   path.Int64Path
	avgRowLength                path.Int64Path
	dataLength                  path.Int64Path
	maxDataLength               path.Int64Path
	indexLength                 path.Int64Path
	dataFree                    path.Int64Path
	createTime                  path.TimePath
	updateTime                  path.TimePath
	checkTime                   path.TimePath
	checksum                    path.Int64Path
	partitionComment            path.StringPath
	nodegroup                   path.StringPath
	tablespaceName              path.StringPath
}

// core.Table Functions

func (q QPartitions) GetColumns() []core.Column {
	return []core.Column{
		q.tableCatalog,
		q.tableSchema,
		q.tableName,
		q.partitionName,
		q.subpartitionName,
		q.partitionOrdinalPosition,
		q.subpartitionOrdinalPosition,
		q.partitionMethod,
		q.subpartitionMethod,
		q.partitionExpression,
		q.subpartitionExpression,
		q.partitionDescription,
		q.tableRows,
		q.avgRowLength,
		q.dataLength,
		q.maxDataLength,
		q.indexLength,
		q.dataFree,
		q.createTime,
		q.updateTime,
		q.checkTime,
		q.checksum,
		q.partitionComment,
		q.nodegroup,
		q.tablespaceName,
	}
}

func (q QPartitions) GetSQL(d core.Dialect, sql core.SQL) error {
	return path.ExpandTableWithDialect(d, q, sql)
}

func (q QPartitions) GetAlias() string {
	return q._alias
}

func (q QPartitions) GetName() string {
	return "PARTITIONS"
}

func (q QPartitions) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QPartitions) TableCatalog() path.StringPath {
	return q.tableCatalog
}

func (q QPartitions) TableSchema() path.StringPath {
	return q.tableSchema
}

func (q QPartitions) TableName() path.StringPath {
	return q.tableName
}

func (q QPartitions) PartitionName() path.StringPath {
	return q.partitionName
}

func (q QPartitions) SubpartitionName() path.StringPath {
	return q.subpartitionName
}

func (q QPartitions) PartitionOrdinalPosition() path.Int64Path {
	return q.partitionOrdinalPosition
}

func (q QPartitions) SubpartitionOrdinalPosition() path.Int64Path {
	return q.subpartitionOrdinalPosition
}

func (q QPartitions) PartitionMethod() path.StringPath {
	return q.partitionMethod
}

func (q QPartitions) SubpartitionMethod() path.StringPath {
	return q.subpartitionMethod
}

func (q QPartitions) PartitionExpression() path.StringPath {
	return q.partitionExpression
}

func (q QPartitions) SubpartitionExpression() path.StringPath {
	return q.subpartitionExpression
}

func (q QPartitions) PartitionDescription() path.StringPath {
	return q.partitionDescription
}

func (q QPartitions) TableRows() path.Int64Path {
	return q.tableRows
}

func (q QPartitions) AvgRowLength() path.Int64Path {
	return q.avgRowLength
}

func (q QPartitions) DataLength() path.Int64Path {
	return q.dataLength
}

func (q QPartitions) MaxDataLength() path.Int64Path {
	return q.maxDataLength
}

func (q QPartitions) IndexLength() path.Int64Path {
	return q.indexLength
}

func (q QPartitions) DataFree() path.Int64Path {
	return q.dataFree
}

func (q QPartitions) CreateTime() path.TimePath {
	return q.createTime
}

func (q QPartitions) UpdateTime() path.TimePath {
	return q.updateTime
}

func (q QPartitions) CheckTime() path.TimePath {
	return q.checkTime
}

func (q QPartitions) Checksum() path.Int64Path {
	return q.checksum
}

func (q QPartitions) PartitionComment() path.StringPath {
	return q.partitionComment
}

func (q QPartitions) Nodegroup() path.StringPath {
	return q.nodegroup
}

func (q QPartitions) TablespaceName() path.StringPath {
	return q.tablespaceName
}
