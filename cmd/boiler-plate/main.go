package main

import (
	"net/http"

	"github.com/Perwira-Media/go-boilerplate-backend/config"
	"github.com/Perwira-Media/go-boilerplate-backend/databases/postgres"
	handler "github.com/Perwira-Media/go-boilerplate-backend/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	configFileLocation = "conf/boiler-config.yaml"
	driverPostgres     = "postgres"
)

// Handlers is the interface for handler object
type Handlers interface {
	Welcome(w http.ResponseWriter, r *http.Request)
}

// NewRouter initiates new router object
func NewRouter(handler Handlers) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Welcome).Methods(http.MethodGet)

	return r
}

func main() {
	logrus.Println("Perwira Media Boiler API Service..")
	config, err := config.GetConfig(configFileLocation)
	if err != nil {
		logrus.Fatalf("Unable to load configuration: %v", err)
	}

	postgresdb, err := postgres.NewPostgresDB(driverPostgres, config.SourceData.PostgresHost, config.SourceData.PostgresUsername,
		config.SourceData.PostgresPassword, config.SourceData.PostgresDBName, config.SourceData.PostgresTimeout, config.SourceData.PostgresPort)
	if err != nil {
		logrus.Errorf("Cannot connect to postgresdb")
		return
	}
	serviceHandler := handler.NewHandler(postgresdb)

	r := NewRouter(serviceHandler)

	logrus.Printf("Starting http server at %v", config.ServiceData.Address)
	err = http.ListenAndServe(config.ServiceData.Address, r)
	if err != nil {
		logrus.Fatalf("Unable to run http server: %v", err)
	}
	logrus.Println("Stopping API Service...")
}
