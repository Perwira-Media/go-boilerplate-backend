package postgres_test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Perwira-Media/go-boilerplate-backend/databases/postgres"
)

func TestGetAllData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	postgresdb := postgres.Postgres{
		DB: db,
	}

	t.Run("TestGetAllData_OK", func(t *testing.T) {
		returnRows := sqlmock.NewRows([]string{"name"}).
			AddRow("test1")

		mock.ExpectQuery("SELECT").WillReturnRows(returnRows)

		_, err = postgresdb.GetAllData()
		if err != nil {
			t.Fatalf("it should not be error, but have: %v", err)
		}
	})

	t.Run("TestGetAllData_NOK Scan", func(t *testing.T) {
		returnRows := sqlmock.NewRows([]string{"name"}).
			AddRow(nil)

		mock.ExpectQuery("SELECT").WillReturnRows(returnRows)

		_, err = postgresdb.GetAllData()
		if err == nil {
			t.Fatalf("it should be error")
		}
	})

	t.Run("TestGetAllData_NOK Query", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("Error intentionally"))

		_, err = postgresdb.GetAllData()
		if err == nil {
			t.Fatalf("it should be error")
		}
	})
}
