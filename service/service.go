package service

import (
	"fmt"
	"log"
	"os"

	"github.com/adantop/golang-bootcamp-2020/db"
	"github.com/adantop/golang-bootcamp-2020/fs"
	"github.com/adantop/golang-bootcamp-2020/repo"
)

// Service is the pokemon service that can be used to get the data
type Service struct {
	DS repo.DataSource
}

// New creates a new service
func New(svcType string) (Service, error) {
	var (
		ds  repo.DataSource
		svc = Service{ds}
	)

	switch svcType {
	case "sqlite3":
		log.Println("Using SQLite3")
		dbfile, err := getEnv("SQLITE3_DBFILE")

		if err != nil {
			return svc, err
		}

		err = db.OpenSQLite3(dbfile, &svc.DS)
		return svc, err

	case "postgres":
		log.Println("Using PostgreSQL")
		pgDSN, err := getEnv("PG_DSN")
		if err != nil {
			return svc, err
		}

		err = db.OpenPostgreSQL(pgDSN, &svc.DS)
		return svc, err

	default:
		log.Println("Using default source: CSV")
		csvfile, err := getEnv("CSV_FILE")
		if err != nil {
			return svc, err
		}

		err = fs.OpenCSV(csvfile, &svc.DS)
		return svc, err

	}
}

func getEnv(varName string) (value string, err error) {
	value = os.Getenv(varName)

	if value == "" {
		err = fmt.Errorf("Environment variable %s not set", varName)
	}
	return
}
