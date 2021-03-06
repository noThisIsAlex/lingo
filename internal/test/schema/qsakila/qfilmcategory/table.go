// Code generated by Lingo for table sakila.film_category - DO NOT EDIT

package qfilmcategory

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QFilmCategory {
	return newQFilmCategory(alias)
}

func New() QFilmCategory {
	return newQFilmCategory("")
}

func newQFilmCategory(alias string) QFilmCategory {
	q := QFilmCategory{_alias: alias}
	q.filmId = path.NewInt16(q, "film_id")
	q.categoryId = path.NewInt8(q, "category_id")
	q.lastUpdate = path.NewTime(q, "last_update")
	return q
}

type QFilmCategory struct {
	_alias     string
	filmId     path.Int16
	categoryId path.Int8
	lastUpdate path.Time
}

// core.Table Functions

func (q QFilmCategory) GetColumns() []core.Column {
	return []core.Column{
		q.filmId,
		q.categoryId,
		q.lastUpdate,
	}
}

func (q QFilmCategory) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QFilmCategory) GetAlias() string {
	return q._alias
}

func (q QFilmCategory) GetName() string {
	return "film_category"
}

func (q QFilmCategory) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QFilmCategory) FilmId() path.Int16 {
	return q.filmId
}

func (q QFilmCategory) CategoryId() path.Int8 {
	return q.categoryId
}

func (q QFilmCategory) LastUpdate() path.Time {
	return q.lastUpdate
}
