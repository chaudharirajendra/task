package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/heroku/task/lib/db"

	"github.com/heroku/task/handler"

	"github.com/gorilla/mux"
)

// main function to boot up everything
func main() {

	router := mux.NewRouter()

	if err := db.Configure(); err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	//Routes
	router.HandleFunc("/moviebytitle/{title}", handler.HandleGetMovieByTitle).Methods("GET")

	router.HandleFunc("/movie/{id}", handler.HandleGetMovieByID).Methods("GET")

	router.HandleFunc("/moviesbyyear", handler.HandleMoviesByYear).Methods("GET")

	router.HandleFunc("/moviesbyrating", handler.HandleMoviesByRating).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
