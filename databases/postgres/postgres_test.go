package postgres_test

import (
	"testing"

	"github.com/Perwira-Media/go-boilerplate-backend/databases/postgres"
)

func TestNewPostgresDB(t *testing.T) {
	driver := "postgres"
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "your-password"
	dbname := "calhounio_demo"
	timeout := 5

	t.Run("Unknown driver OK", func(t *testing.T) {
		_, err := postgres.NewPostgresDB(driver, host, user, password, dbname, timeout, port)
		if err != nil {
			t.Fatalf("Create connection should not be error but found: %v", err)
		}
	})

	t.Run("Unknown driver NOK", func(t *testing.T) {
		driver = "postgre"
		_, err := postgres.NewPostgresDB(driver, host, user, password, dbname, timeout, port)
		if err == nil {
			t.Fatal("Create connection should be error")
		}
	})
}
