package router

import (
	"encoding/json"
	"net/http"
)

func initializeRoutes(router *http.ServeMux) {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		json.NewEncoder(w).Encode(JSONTest{
			Message: "pong",
		})
	})
}
