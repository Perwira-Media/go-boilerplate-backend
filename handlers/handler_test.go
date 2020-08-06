package handlers_test

import (
	"fmt"
	"testing"

	handler "github.com/Perwira-Media/go-boilerplate-backend/handlers"
)

type mockedPostgres struct {
	err          map[string]bool
	errStatement error
}

func TestNewHandler(t *testing.T) {
	mongo := &mockedPostgres{
		err:          map[string]bool{},
		errStatement: fmt.Errorf("Intentionally Error."),
	}

	t.Run("OK", func(t *testing.T) {
		testhandler := handler.NewHandler(mongo)
		if testhandler == nil {
			t.Fatalf("Handler should not be nil")
		}
	})
}
