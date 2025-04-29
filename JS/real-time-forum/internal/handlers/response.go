package handlers

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, code int, resp any) {
	response, _ := json.Marshal(resp)
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}