package model

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code,
		Response{
			Code:    code,
			Status:  "Error",
			Message: message,
		})
}

func RespondWithErrors(w http.ResponseWriter, code int, message string, errors []string) {
	RespondWithJSON(w, code, Response{
		Code:    code,
		Status:  "Error",
		Message: message,
		Errors:  errors,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
