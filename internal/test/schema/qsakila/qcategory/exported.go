// Code generated by Lingo for table sakila.category - DO NOT EDIT

package qcategory

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QCategory {
	return instance
}

func CategoryId() path.Int8Path {
	return instance.categoryId
}

func Name() path.StringPath {
	return instance.name
}

func LastUpdate() path.TimePath {
	return instance.lastUpdate
}