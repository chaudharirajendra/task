package db

import (
	"log"

	// sql driver

	"github.com/jmoiron/sqlx"
)

// DB is the global database connection pool
var DB *sqlx.DB

// Configure sets the domain configuration
func Configure() error {

	var err error

	//DB, err = sqlx.Connect("mysql", "username:password@(host:port)/databasename")

	DB, err = sqlx.Connect("mysql", "root:1234@(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
