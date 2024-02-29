package handler

import (
	"encoding/json"
	"net/http"
)

type JSONTest struct {
	Message string `json:"message"`
}

const jsonContentType = "application/json"

func GetOpening(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(JSONTest{
		Message: "GET opening",
	})
}

func GetAllOpenings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(JSONTest{
		Message: "GET openings",
	})
}

func CreateOpening(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(JSONTest{
		Message: "POST opening",
	})
}

func DeleteOpening(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(JSONTest{
		Message: "DELETE opening",
	})
}

func UpdateOpening(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(JSONTest{
		Message: "PUT opening",
	})
}
