// Code generated by Lingo for table information_schema.CHARACTER_SETS - DO NOT EDIT

package qcharactersets

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QCharacterSets {
	return instance
}

func CharacterSetName() path.StringPath {
	return instance.characterSetName
}

func DefaultCollateName() path.StringPath {
	return instance.defaultCollateName
}

func Description() path.StringPath {
	return instance.description
}

func Maxlen() path.Int64Path {
	return instance.maxlen
}
