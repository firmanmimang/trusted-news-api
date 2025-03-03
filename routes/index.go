package routes

import "github.com/gorilla/mux"

func IndexRoutes(r *mux.Router) {
	// API ROUTE
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	NewsRoutes(apiV1)
}
