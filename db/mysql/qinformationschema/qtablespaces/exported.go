// Code generated by Lingo for table information_schema.TABLESPACES - DO NOT EDIT

package qtablespaces

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QTablespaces {
	return instance
}

func TablespaceName() path.StringPath {
	return instance.tablespaceName
}

func Engine() path.StringPath {
	return instance.engine
}

func TablespaceType() path.StringPath {
	return instance.tablespaceType
}

func LogfileGroupName() path.StringPath {
	return instance.logfileGroupName
}

func ExtentSize() path.Int64Path {
	return instance.extentSize
}

func AutoextendSize() path.Int64Path {
	return instance.autoextendSize
}

func MaximumSize() path.Int64Path {
	return instance.maximumSize
}

func NodegroupId() path.Int64Path {
	return instance.nodegroupId
}

func TablespaceComment() path.StringPath {
	return instance.tablespaceComment
}
