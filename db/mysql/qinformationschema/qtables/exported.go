// Code generated by Lingo for table information_schema.TABLES - DO NOT EDIT

package qtables

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QTables {
	return instance
}

func TableCatalog() path.StringPath {
	return instance.tableCatalog
}

func TableSchema() path.StringPath {
	return instance.tableSchema
}

func TableName() path.StringPath {
	return instance.tableName
}

func TableType() path.StringPath {
	return instance.tableType
}

func Engine() path.StringPath {
	return instance.engine
}

func Version() path.Int64Path {
	return instance.version
}

func RowFormat() path.StringPath {
	return instance.rowFormat
}

func TableRows() path.Int64Path {
	return instance.tableRows
}

func AvgRowLength() path.Int64Path {
	return instance.avgRowLength
}

func DataLength() path.Int64Path {
	return instance.dataLength
}

func MaxDataLength() path.Int64Path {
	return instance.maxDataLength
}

func IndexLength() path.Int64Path {
	return instance.indexLength
}

func DataFree() path.Int64Path {
	return instance.dataFree
}

func AutoIncrement() path.Int64Path {
	return instance.autoIncrement
}

func CreateTime() path.TimePath {
	return instance.createTime
}

func UpdateTime() path.TimePath {
	return instance.updateTime
}

func CheckTime() path.TimePath {
	return instance.checkTime
}

func TableCollation() path.StringPath {
	return instance.tableCollation
}

func Checksum() path.Int64Path {
	return instance.checksum
}

func CreateOptions() path.StringPath {
	return instance.createOptions
}

func TableComment() path.StringPath {
	return instance.tableComment
}