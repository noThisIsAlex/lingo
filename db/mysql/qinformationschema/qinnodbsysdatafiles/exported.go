// Code generated by Lingo for table information_schema.INNODB_SYS_DATAFILES - DO NOT EDIT

package qinnodbsysdatafiles

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QInnodbSysDatafiles {
	return instance
}

func Space() path.IntPath {
	return instance.space
}

func Path() path.StringPath {
	return instance.path
}