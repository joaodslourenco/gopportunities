package router

import (
	"log"
	"net/http"
)

func Initialize() {
	router := http.NewServeMux()

	initializeRoutes(router)

	log.Fatal(http.ListenAndServe(":5000", router))

}
