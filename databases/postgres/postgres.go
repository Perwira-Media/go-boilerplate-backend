package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Postgres is the postgres object
type Postgres struct {
	DB *sql.DB
}

const (
	dataCollections = "data"
)

// NewPostgresDB creates new postgresdb instance and returns Postgres object
func NewPostgresDB(driver, host, username, password, dbname string, timeout, port int) (*Postgres, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d sslmode=disable", host, port, username, password, dbname, timeout)
	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return nil, err
	}

	postgresDB := &Postgres{
		DB: db,
	}

	return postgresDB, nil
}
