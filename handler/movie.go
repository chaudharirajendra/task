package handler

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heroku/task/domain/movie"
)

//HandleUpdateMovieGenres updates the movie according to genres
func HandleUpdateMovieGenres(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	ID := params["id"]
	rating := params["genres"]

	genres, err := movie.UpdateMovieGenres(ID, rating)
	if handleError(w, r, ServError, err) {
		return
	}

	send(w, r, &genres)

}

//HandleUpdateMovieRating updates the movie according to rating
func HandleUpdateMovieRating(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	ID := params["id"]
	rating := params["rating"]

	movie, err := movie.UpdateMovieRating(ID, rating)
	if handleError(w, r, ServError, err) {
		return
	}

	send(w, r, &movie)

}

//HandleMoviesByGenres return array of movies object based on genres
func HandleMoviesByGenres(w http.ResponseWriter, r *http.Request) {

	genres := r.URL.Query().Get("genres")
	movies, err := movie.GetMoviesByGenres(genres)
	if handleError(w, r, ServError, err) {
		return
	}

	send(w, r, &movies)
}

//HandleMoviesByRating returns array movie object based on rating
func HandleMoviesByRating(w http.ResponseWriter, r *http.Request) {

	rating := r.URL.Query().Get("rating")
	movies, err := movie.GetMoviesByRating(rating)
	if handleError(w, r, ServError, err) {
		return
	}

	send(w, r, &movies)
}

//HandleMoviesByYear returns array movie object based on year
func HandleMoviesByYear(w http.ResponseWriter, r *http.Request) {

	year := r.URL.Query().Get("year")
	movies, err := movie.GetMoviesByYear(year)
	if handleError(w, r, ServError, err) {
		return
	}

	send(w, r, &movies)
}

//HandleGetMovieByID returns movie object by Id
func HandleGetMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID := params["id"]

	movieObj, err := movie.GetMovieByID(ID)
	if handleError(w, r, ServError, err) {
		return
	}

	send(w, r, &movieObj)
}

// HandleGetMovieByTitle returns an movie object if provided text match with movie column
func HandleGetMovieByTitle(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	title := params["title"]

	if title == "" {
		handleError(w, r, BadRequest, errors.New("title required"))
		return
	}

	movieObj, err := movie.GetMovieByTitle(title)
	if handleError(w, r, ServError, err) {
		return
	}

	send(w, r, &movieObj)
}
