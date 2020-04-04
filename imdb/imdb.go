package imdb

import (
	"github.com/fatih/color"
	"github.com/kenshaw/imdb"
)

//Film parametre adı olarak aldığı film adını imdb listesinde arayıp sonucu döndürüyor
func Film(filmName string) (string, string, string, string, string, string) {
	cl := imdb.New("ae609f1a")
	res, err := cl.MovieByTitle(filmName, "")
	if err != nil {
		color.Red(err.Error())
	}
	return res.Title, res.Actors, res.Year, res.ImdbRating, res.Poster, res.Error
}
