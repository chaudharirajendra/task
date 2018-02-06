package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/heroku/task/lib/db"

	"github.com/heroku/task/handler"

	"github.com/gorilla/mux"
)

// main function to boot up everything
func main() {

	router := mux.NewRouter()

	// database
	dbconf := db.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	if err := db.Configure(&dbconf); err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	//Routes
	router.HandleFunc("/moviebytitle/{title}", handler.HandleGetMovieByTitle).Methods("GET")

	router.HandleFunc("/movie/{id}", handler.HandleGetMovieByID).Methods("GET")

	router.HandleFunc("/moviesbyyear", handler.HandleMoviesByYear).Methods("GET")

	router.HandleFunc("/moviesbyrating", handler.HandleMoviesByRating).Methods("GET")

	router.HandleFunc("/moviesbygenres", handler.HandleMoviesByGenres).Methods("GET")

	router.HandleFunc("/updatemovierating/{id}/{rating}", handler.HandleUpdateMovieRating).Methods("PUT")

	router.HandleFunc("/updatemoviegenres/{id}/{genres}", handler.HandleUpdateMovieGenres).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}
