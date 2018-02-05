package db

import (
	"fmt"
	"log"

	// sql driver

	"github.com/jmoiron/sqlx"
)

// Config is the config items for the domain
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var localConfig *Config

// DB is the global database connection pool
var DB *sqlx.DB

// Configure sets the domain configuration
func Configure(c *Config) error {

	localConfig = c

	var err error

	dsn := fmt.Sprintf("%v:%v@(%v:%v)/%v%v", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, "?parseTime=true")

	//DB, err = sqlx.Connect("mysql", "username:password@(host:port)/databasename")

	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
