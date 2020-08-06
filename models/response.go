package models

// Response used for sending response
type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Status  interface{} `json:"status,omitempty"`
}
