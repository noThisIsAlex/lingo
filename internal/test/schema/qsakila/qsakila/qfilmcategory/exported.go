// Code generated by Lingo for table sakila.film_category - DO NOT EDIT

package qfilmcategory

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QFilmCategory {
	return instance
}

func FilmId() path.Int16Path {
	return instance.filmId
}

func CategoryId() path.BoolPath {
	return instance.categoryId
}

func LastUpdate() path.TimePath {
	return instance.lastUpdate
}
