// Code generated by Lingo for schema sakila - DO NOT EDIT

package qsakila

var instance = schema{_name: "sakila"}

func GetSchema() *schema {
	return &instance
}

type schema struct {
	_name string
}

func (s schema) GetName() string {
	return s._name
}
