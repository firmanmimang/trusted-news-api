package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/firmanmimang/go-api-trusted-news/config"
	"github.com/firmanmimang/go-api-trusted-news/middlewares"
	"github.com/firmanmimang/go-api-trusted-news/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	r.Use(middlewares.WithDB(config.DB))

	// API ROUTE
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	routes.NewsRoutes(apiV1)

	log.Println("server running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
