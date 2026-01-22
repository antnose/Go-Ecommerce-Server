package cmd

import (
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	fmt.Println("Server running on port http://localhost:3001")

	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
