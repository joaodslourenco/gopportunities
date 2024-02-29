package router

import (
	"fmt"
	"net/http"

	"github.com/joaodslourenco/gopportunities/handler"
)

func generateEndpoint(method, endpoint string) string {
	return fmt.Sprintf("%s %s", method, endpoint)
}

func initializeRoutes(router *http.ServeMux) {
	router.HandleFunc(generateEndpoint(http.MethodGet, "/opening"), handler.GetOpening)
	router.HandleFunc(generateEndpoint(http.MethodPost, "/opening"), handler.CreateOpening)
	router.HandleFunc(generateEndpoint(http.MethodDelete, "/opening"), handler.DeleteOpening)
	router.HandleFunc(generateEndpoint(http.MethodPut, "/opening"), handler.UpdateOpening)
	router.HandleFunc(generateEndpoint(http.MethodGet, "/openings"), handler.GetAllOpenings)
}
