package router

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONTest struct {
	Message string `json:"message"`
}

const jsonContentType = "application/json"

func Initialize() {
	router := http.NewServeMux()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		json.NewEncoder(w).Encode(JSONTest{
			Message: "pong"})
	})
	log.Fatal(http.ListenAndServe(":5000", router))

}
