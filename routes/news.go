package routes

import (
	"github.com/firmanmimang/go-api-trusted-news/controllers/newscontroller"
	"github.com/gorilla/mux"
)

func NewsRoutes(r *mux.Router) {
	r.HandleFunc("/news", newscontroller.Index).Methods("GET")
}
