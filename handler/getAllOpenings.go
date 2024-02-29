package handler

import (
	"encoding/json"
	"net/http"
)

func GetAllOpenings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(JSONTest{
		Message: "GET openings",
	})
}
