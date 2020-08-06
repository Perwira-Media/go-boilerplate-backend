package handlers

import (
	model "github.com/Perwira-Media/go-boilerplate-backend/models"
)

// Handler is used for initiating handler object
type Handler struct {
	postgresdb Postgres
}

// Postgres is the interface for database object
type Postgres interface {
	GetAllData() ([]model.Data, error)
}

// NewHandler create a new handler and returns handler object
func NewHandler(postgres Postgres) *Handler {
	return &Handler{
		postgresdb: postgres,
	}
}
