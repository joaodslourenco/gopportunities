package router

import (
	"log"
	"net/http"
)

type JSONTest struct {
	Message string `json:"message"`
}

const jsonContentType = "application/json"

func Initialize() {
	router := http.NewServeMux()

	initializeRoutes(router)

	log.Fatal(http.ListenAndServe(":5000", router))

}
