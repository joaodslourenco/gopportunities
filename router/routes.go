package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func generateEndpoint(method, endpoint string) string {
	return fmt.Sprintf("%s %s", method, endpoint)
}

func initializeRoutes(router *http.ServeMux) {

	router.HandleFunc(generateEndpoint(http.MethodGet, "/opening"), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		json.NewEncoder(w).Encode(JSONTest{
			Message: "GET opening",
		})
	})

	router.HandleFunc(generateEndpoint(http.MethodPost, "/opening"), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		json.NewEncoder(w).Encode(JSONTest{
			Message: "POST opening",
		})
	})

	router.HandleFunc(generateEndpoint(http.MethodDelete, "/opening"), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		json.NewEncoder(w).Encode(JSONTest{
			Message: "DELETE opening",
		})
	})

	router.HandleFunc(generateEndpoint(http.MethodPut, "/opening"), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		json.NewEncoder(w).Encode(JSONTest{
			Message: "PUT opening",
		})
	})

	router.HandleFunc(generateEndpoint(http.MethodGet, "/openings"), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		json.NewEncoder(w).Encode(JSONTest{
			Message: "GET openings",
		})
	})

}
