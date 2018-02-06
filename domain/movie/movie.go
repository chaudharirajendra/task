package movie

import (
	"fmt"

	"github.com/eefret/gomdb"
	"github.com/pkg/errors"

	"github.com/heroku/task/lib/db"
)

// The Movie Type (more like an object)
type Movie struct {
	ID           int64        `db:"id" json:"id"`
	Title        string       `db:"title" json:"title"`
	ReleasedYear string       `db:"released_year" json:"released_year"`
	Rating       db.NullInt64 `db:"rating" json:"rating"`
	Generes      []Generes    `json:"generes"`
}

//The Generes Type (subobject of Movie)
type Generes struct {
	ID          int64  `db:"id" json:"id"`
	MovieID     int64  `db:"movie_id" json:"movie_id"`
	GeneresName string `db:"genres_name" json:"genres_name"`
}

//UpdateMovieGenres updates the genres of the movie.
func UpdateMovieGenres(id, genresname string) (Generes, error) {

	genres := Generes{}

	sqlstr := `update genres set genres_name='` + genresname + `' where id=` + id + ``

	_, err := db.DB.Exec(sqlstr)

	if err != nil {
		return genres, err
	}

	genres, err = GetGenresByID(id)

	if err != nil {
		return genres, err
	}

	return genres, nil
}

//UpdateMovieRating updates the rating of the movie.
func UpdateMovieRating(id, rating string) (Movie, error) {

	movie := Movie{}

	sqlstr := `update movie set rating=` + rating + ` where id=` + id + ``

	_, err := db.DB.Exec(sqlstr)
	if err != nil {
		return movie, err
	}

	movie, err = GetMovieByID(id)

	if err != nil {
		return movie, err
	}

	return movie, nil
}

//GetMoviesByGenres returns the array of movie object based on genres.
func GetMoviesByGenres(genres string) ([]Movie, error) {

	movies := []Movie{}

	sqlstr := `SELECT m.id,m.title,m.released_year,m.rating from movie m LEFT JOIN genres g ON m.id=g.movie_id where genres_name LIKE   '%` + genres + `%' `

	err := db.DB.Select(&movies, sqlstr)
	if err != nil {
		return movies, err
	}

	for i, movie := range movies {

		genrs := []Generes{}

		genrs, err := GetAllGeneres(fmt.Sprint(movie.ID))

		if err != nil {
			return movies, err
		}

		movies[i].Generes = genrs

	}

	return movies, nil

}

//GetMoviesByRating returns the array of movie object based on rating.
func GetMoviesByRating(rating string) ([]Movie, error) {

	movies := []Movie{}

	sqlstr := `SELECT id,title,released_year,rating from movie where rating > '` + rating + `' `

	err := db.DB.Select(&movies, sqlstr)
	if err != nil {
		return movies, err
	}

	for i, movie := range movies {

		genrs := []Generes{}

		genrs, err := GetAllGeneres(fmt.Sprint(movie.ID))

		if err != nil {
			return movies, err
		}

		movies[i].Generes = genrs

	}

	return movies, nil
}

//GetMoviesByYear returns the array of movie object based on year.
func GetMoviesByYear(year string) ([]Movie, error) {

	movies := []Movie{}

	sqlstr := `SELECT id,title,released_year,rating from movie where released_year='` + year + `' `

	fmt.Println(sqlstr)
	err := db.DB.Select(&movies, sqlstr)
	if err != nil {
		return movies, err
	}

	for i, movie := range movies {

		genrs := []Generes{}

		genrs, err := GetAllGeneres(fmt.Sprint(movie.ID))

		if err != nil {
			return movies, err
		}

		movies[i].Generes = genrs

	}

	return movies, nil
}

//GetMovieByTitle returns the movie by title.
func GetMovieByTitle(title string) (Movie, error) {

	movie := Movie{}

	//check movie with title is available or not in our database.
	if IsMovieWithTitleAvailable(title) {

		sqlstr := `SELECT id,title,released_year,rating from movie where title='` + title + `'  order by id desc limit 1`
		err := db.DB.Get(&movie, sqlstr)
		if err != nil {
			return movie, err
		}
	} else {

		//if the data isn't available in our database then call imdb to get the results.
		api := gomdb.Init("83c448c5")
		query := &gomdb.QueryData{Title: title, SearchType: gomdb.MovieSearch}
		res, err := api.Search(query)
		if err != nil {
			fmt.Println(err)
			return movie, err
		}

		var mID int64

		for _, resp := range res.Search {

			sqlstr := `
			INSERT INTO movie
			(title, released_year)
			VALUES
			(?,?)`

			res, err := db.DB.Exec(sqlstr, resp.Title, resp.Year)
			if err != nil {
				return movie, err
			}
			mID, _ = res.LastInsertId()

		}

		movie, err = GetMovieByID(fmt.Sprint(mID))

		if err != nil {
			return movie, err
		}
	}

	return movie, nil
}

//GetGenresByID returns the genres by id.
func GetGenresByID(id string) (Generes, error) {

	genres := Generes{}

	if CheckIfGenresFind(id) {
		sqlstr := `SELECT id,movie_id,genres_name from genres where id=` + id + ` `
		err := db.DB.Get(&genres, sqlstr)

		if err != nil {
			return genres, err
		}
	} else {
		return genres, errors.New("Genres not Found for ID:" + id)
	}

	return genres, nil
}

//GetMovieByID returns the movie by id.
func GetMovieByID(mid string) (Movie, error) {

	movie := Movie{}

	if CheckIfMovieFind(mid) {
		sqlstr := `SELECT id,title,released_year,rating from movie where id=` + mid + ` `
		err := db.DB.Get(&movie, sqlstr)
		fmt.Println(err)
		if err != nil {
			return movie, err
		}
	} else {
		return movie, errors.New("Movie not Found for ID:" + mid)
	}

	gerres, err := GetAllGeneres(mid)

	if err != nil {
		return movie, err
	}

	movie.Generes = gerres

	return movie, nil
}

//GetAllGeneres return all the genres based on movieid
func GetAllGeneres(mid string) ([]Generes, error) {

	genres := []Generes{}

	sqlstr := `select id,movie_id,genres_name from genres where movie_id=` + mid + ``
	err := db.DB.Select(&genres, sqlstr)
	if err != nil {
		return genres, err
	}

	return genres, nil
}

//IsMovieWithTitleAvailable returns true if record found otherwise false
func IsMovieWithTitleAvailable(title string) bool {

	var cnt int64
	sqlstr := `SELECT count(*) from movie where title='` + title + `' `

	err := db.DB.Get(&cnt, sqlstr)
	fmt.Println(err)
	if err != nil {
		return false
	}

	if cnt > 0 {
		return true
	}
	return false
}

//CheckIfMovieFind returns true if record found otherwise false
func CheckIfMovieFind(mid string) bool {

	var cnt int64
	sqlstr := `SELECT count(*) from movie where id=` + mid + ` `

	err := db.DB.Get(&cnt, sqlstr)

	if err != nil {
		return false
	}

	if cnt > 0 {
		return true
	}
	return false
}

//CheckIfGenresFind returns true if record found otherwise false
func CheckIfGenresFind(id string) bool {

	var cnt int64
	sqlstr := `SELECT count(*) from genres where id=` + id + ` `

	err := db.DB.Get(&cnt, sqlstr)

	if err != nil {
		return false
	}

	if cnt > 0 {
		return true
	}
	return false
}
