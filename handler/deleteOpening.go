package handler

import (
	"encoding/json"
	"net/http"
)

func DeleteOpening(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(JSONTest{
		Message: "DELETE opening",
	})
}
