package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const jsonContentType = "application/json"

type JSONTest struct {
	Message string `json:"message"`
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		json.NewEncoder(w).Encode(JSONTest{
			Message: "ping"})
	})

	log.Fatal(http.ListenAndServe(":5000", router))
}
