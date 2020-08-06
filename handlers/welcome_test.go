package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	handler "github.com/Perwira-Media/go-boilerplate-backend/handlers"
	model "github.com/Perwira-Media/go-boilerplate-backend/models"
	"github.com/gorilla/mux"
)

func (db *mockedPostgres) GetAllData() ([]model.Data, error) {
	if db.err["GetAllData"] {
		return nil, db.errStatement
	}

	return []model.Data{
		model.Data{
			Name: "test1",
		},
		model.Data{
			Name: "test2",
		},
	}, nil
}

func TestWelcome(t *testing.T) {
	mongo := &mockedPostgres{
		err:          map[string]bool{},
		errStatement: fmt.Errorf("Intentionally Error."),
	}

	testhandler := handler.NewHandler(mongo)
	r := mux.NewRouter()

	r.HandleFunc("/", testhandler.Welcome).Methods("GET")

	httpserver := httptest.NewServer(r)
	defer httpserver.Close()

	serverURL, _ := url.Parse(httpserver.URL)

	targetPath := fmt.Sprintf("%v", serverURL)
	req, _ := http.NewRequest("GET", targetPath, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Unable to get worker status: %v", err)
	}
	defer resp.Body.Close()

	t.Run("OK", func(t *testing.T) {
		req, _ := http.NewRequest("GET", targetPath, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Unable to get worker status: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Status code should be %v, not %v", http.StatusOK, resp.StatusCode)
		}
	})
}
