// Code generated by Lingo for table information_schema.KEY_COLUMN_USAGE - DO NOT EDIT

package qkeycolumnusage

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QKeyColumnUsage {
	return instance
}

func ConstraintCatalog() path.StringPath {
	return instance.constraintCatalog
}

func ConstraintSchema() path.StringPath {
	return instance.constraintSchema
}

func ConstraintName() path.StringPath {
	return instance.constraintName
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

func ColumnName() path.StringPath {
	return instance.columnName
}

func OrdinalPosition() path.Int64Path {
	return instance.ordinalPosition
}

func PositionInUniqueConstraint() path.Int64Path {
	return instance.positionInUniqueConstraint
}

func ReferencedTableSchema() path.StringPath {
	return instance.referencedTableSchema
}

func ReferencedTableName() path.StringPath {
	return instance.referencedTableName
}

func ReferencedColumnName() path.StringPath {
	return instance.referencedColumnName
}
