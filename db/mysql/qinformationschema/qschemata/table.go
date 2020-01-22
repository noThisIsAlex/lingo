// Code generated by Lingo for table information_schema.SCHEMATA - DO NOT EDIT

package qschemata

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QSchemata {
	return newQSchemata(alias)
}

func New() QSchemata {
	return newQSchemata("")
}

func newQSchemata(alias string) QSchemata {
	q := QSchemata{_alias: alias}
	q.catalogName = path.NewStringPath(q, "CATALOG_NAME")
	q.schemaName = path.NewStringPath(q, "SCHEMA_NAME")
	q.defaultCharacterSetName = path.NewStringPath(q, "DEFAULT_CHARACTER_SET_NAME")
	q.defaultCollationName = path.NewStringPath(q, "DEFAULT_COLLATION_NAME")
	q.sqlPath = NewUnknownPathType(q, "SQL_PATH")
	q.defaultEncryption = NewUnknownPathType(q, "DEFAULT_ENCRYPTION")
	return q
}

type QSchemata struct {
	_alias                  string
	catalogName             path.StringPath
	schemaName              path.StringPath
	defaultCharacterSetName path.StringPath
	defaultCollationName    path.StringPath
	sqlPath                 UnknownPathType
	defaultEncryption       UnknownPathType
}

// core.Table Functions

func (q QSchemata) GetColumns() []core.Column {
	return []core.Column{
		q.catalogName,
		q.schemaName,
		q.defaultCharacterSetName,
		q.defaultCollationName,
		q.sqlPath,
		q.defaultEncryption,
	}
}

func (q QSchemata) GetSQL(d core.Dialect, sql core.SQL) error {
	return path.ExpandTableWithDialect(d, q, sql)
}

func (q QSchemata) GetAlias() string {
	return q._alias
}

func (q QSchemata) GetName() string {
	return "SCHEMATA"
}

func (q QSchemata) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QSchemata) CatalogName() path.StringPath {
	return q.catalogName
}

func (q QSchemata) SchemaName() path.StringPath {
	return q.schemaName
}

func (q QSchemata) DefaultCharacterSetName() path.StringPath {
	return q.defaultCharacterSetName
}

func (q QSchemata) DefaultCollationName() path.StringPath {
	return q.defaultCollationName
}

func (q QSchemata) SqlPath() UnknownPathType {
	return q.sqlPath
}

func (q QSchemata) DefaultEncryption() UnknownPathType {
	return q.defaultEncryption
}
