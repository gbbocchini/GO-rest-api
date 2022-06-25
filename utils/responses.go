package utils

import (
	"encoding/json"
	"net/http"
	"rest/models"
)

func SendError(w http.ResponseWriter, statusCode int, err models.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
