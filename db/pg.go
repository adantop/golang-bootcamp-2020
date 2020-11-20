package db

import (
	"database/sql"
	"log"

	"github.com/adantop/golang-bootcamp-2020/repo"
	// Required by database/sql
	_ "github.com/lib/pq"
)

// OpenPostgreSQL overwrites the db module DB var with a pg implementation
func OpenPostgreSQL(dsn string, ds *repo.DataSource) (err error) {
	//var pgDSN = "host=localhost port=5432 user=postgres password=passwd dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Unable to create db object: %v\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not establish connection to the database: %v", err)
	}

	(*ds) = database{db}

	return err
}
