package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use()

	mux := http.NewServeMux()

	// mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts)))

	// mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))

	// mux.Handle("GET /products/{id}", middleware.Logger(http.HandlerFunc(handlers.GetProductById)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))

	fmt.Println("Server running on port http://localhost:3001")

	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
