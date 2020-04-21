// Code generated by Lingo for table sakila.staff - DO NOT EDIT

package qstaff

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QStaff {
	return newQStaff(alias)
}

func New() QStaff {
	return newQStaff("")
}

func newQStaff(alias string) QStaff {
	q := QStaff{_alias: alias}
	q.staffId = path.NewInt8Path(q, "staff_id")
	q.firstName = path.NewStringPath(q, "first_name")
	q.lastName = path.NewStringPath(q, "last_name")
	q.addressId = path.NewInt16Path(q, "address_id")
	q.picture = path.NewUnsupportedPath(q, "picture")
	q.email = path.NewStringPath(q, "email")
	q.storeId = path.NewInt8Path(q, "store_id")
	q.active = path.NewInt8Path(q, "active")
	q.username = path.NewStringPath(q, "username")
	q.password = path.NewStringPath(q, "password")
	q.lastUpdate = path.NewTimePath(q, "last_update")
	return q
}

type QStaff struct {
	_alias     string
	staffId    path.Int8Path
	firstName  path.StringPath
	lastName   path.StringPath
	addressId  path.Int16Path
	picture    path.UnsupportedPath
	email      path.StringPath
	storeId    path.Int8Path
	active     path.Int8Path
	username   path.StringPath
	password   path.StringPath
	lastUpdate path.TimePath
}

// core.Table Functions

func (q QStaff) GetColumns() []core.Column {
	return []core.Column{
		q.staffId,
		q.firstName,
		q.lastName,
		q.addressId,
		q.picture,
		q.email,
		q.storeId,
		q.active,
		q.username,
		q.password,
		q.lastUpdate,
	}
}

func (q QStaff) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QStaff) GetAlias() string {
	return q._alias
}

func (q QStaff) GetName() string {
	return "staff"
}

func (q QStaff) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QStaff) StaffId() path.Int8Path {
	return q.staffId
}

func (q QStaff) FirstName() path.StringPath {
	return q.firstName
}

func (q QStaff) LastName() path.StringPath {
	return q.lastName
}

func (q QStaff) AddressId() path.Int16Path {
	return q.addressId
}

func (q QStaff) Picture() path.UnsupportedPath {
	return q.picture
}

func (q QStaff) Email() path.StringPath {
	return q.email
}

func (q QStaff) StoreId() path.Int8Path {
	return q.storeId
}

func (q QStaff) Active() path.Int8Path {
	return q.active
}

func (q QStaff) Username() path.StringPath {
	return q.username
}

func (q QStaff) Password() path.StringPath {
	return q.password
}

func (q QStaff) LastUpdate() path.TimePath {
	return q.lastUpdate
}