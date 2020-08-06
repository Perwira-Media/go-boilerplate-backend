package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/Perwira-Media/go-boilerplate-backend/models"
)

// Welcome display welcome page
func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	response := model.Response{}

	response.Status = http.StatusOK
	response.Message = "welcome to boiler API"

	jsonedResponse, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonedResponse)
	return
}
