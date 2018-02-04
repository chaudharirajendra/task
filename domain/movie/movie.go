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
}

//GetMoviesByRating returns the array of movie object based on rating.
func GetMoviesByRating(rating string) ([]Movie, error) {

	movies := []Movie{}

	sqlstr := `SELECT id,title,released_year,rating from movie where rating > '` + rating + `' `

	fmt.Println(sqlstr)
	err := db.DB.Select(&movies, sqlstr)
	if err != nil {
		return movies, err
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

	return movie, nil
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

	fmt.Println(err)
	if err != nil {
		return false
	}

	if cnt > 0 {
		return true
	}
	return false
}
