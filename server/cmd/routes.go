package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts)))

	// mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))

	// mux.Handle("GET /products/{id}", middleware.Logger(http.HandlerFunc(handlers.GetProductById)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))
}
