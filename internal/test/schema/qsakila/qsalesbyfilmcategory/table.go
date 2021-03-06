// Code generated by Lingo for table sakila.sales_by_film_category - DO NOT EDIT

package qsalesbyfilmcategory

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QSalesByFilmCategory {
	return newQSalesByFilmCategory(alias)
}

func New() QSalesByFilmCategory {
	return newQSalesByFilmCategory("")
}

func newQSalesByFilmCategory(alias string) QSalesByFilmCategory {
	q := QSalesByFilmCategory{_alias: alias}
	q.category = path.NewString(q, "category")
	q.totalSales = path.NewBinary(q, "total_sales")
	return q
}

type QSalesByFilmCategory struct {
	_alias     string
	category   path.String
	totalSales path.Binary
}

// core.Table Functions

func (q QSalesByFilmCategory) GetColumns() []core.Column {
	return []core.Column{
		q.category,
		q.totalSales,
	}
}

func (q QSalesByFilmCategory) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QSalesByFilmCategory) GetAlias() string {
	return q._alias
}

func (q QSalesByFilmCategory) GetName() string {
	return "sales_by_film_category"
}

func (q QSalesByFilmCategory) GetParent() string {
	return "sakila"
}

// Column Functions

func (q QSalesByFilmCategory) Category() path.String {
	return q.category
}

func (q QSalesByFilmCategory) TotalSales() path.Binary {
	return q.totalSales
}
