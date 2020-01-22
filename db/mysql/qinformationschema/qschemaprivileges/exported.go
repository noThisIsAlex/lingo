// Code generated by Lingo for table information_schema.SCHEMA_PRIVILEGES - DO NOT EDIT

package qschemaprivileges

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QSchemaPrivileges {
	return instance
}

func Grantee() path.StringPath {
	return instance.grantee
}

func TableCatalog() path.StringPath {
	return instance.tableCatalog
}

func TableSchema() path.StringPath {
	return instance.tableSchema
}

func PrivilegeType() path.StringPath {
	return instance.privilegeType
}

func IsGrantable() path.StringPath {
	return instance.isGrantable
}
