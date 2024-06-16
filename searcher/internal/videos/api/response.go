package api

import (
	"encoding/json"
	"net/http"
	"src/internal/word/utils"
)

func videoResponse(data interface{}, w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return
	}
}

func errorResponse(status int, message utils.ErrorResponseStruct, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		return
	}
}
