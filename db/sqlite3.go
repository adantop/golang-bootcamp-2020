package db

import (
	"database/sql"
	"log"

	"github.com/adantop/golang-bootcamp-2020/repo"
	// Registering sqlite3 driver into database/sql
	_ "github.com/mattn/go-sqlite3"
)

// OpenSQLite3 overwrites the db module DB var with an SQLite3 implementation
func OpenSQLite3(dbfile string, ds *repo.DataSource) (err error) {

	db, err := sql.Open("sqlite3", dbfile)

	if err != nil {
		log.Fatalf("Unable to create db object: %v\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not establish connection to the database: %v", err)
	}

	(*ds) = database{db}

	return err
}
