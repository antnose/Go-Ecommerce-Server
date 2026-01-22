package main

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello I am Ibrahim. I Will be software engineer")
}

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/hello", helloHandle) //Old Routing
	// mux.Handle("GET /hello", http.HandlerFunc(helloHandle))
	mux.Handle("GET /hello", global_router.HandleCorsMiddleware(http.HandlerFunc(helloHandle)))

	// mux.HandleFunc("/about", aboutHandler) //Old Routing
	// mux.Handle("GET /about", http.HandlerFunc(aboutHandler))
	mux.Handle("GET /about", global_router.HandleCorsMiddleware(http.HandlerFunc(aboutHandler)))

	// mux.HandleFunc("/products", getProducts) //Old Routing
	// mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("GET /products", global_router.HandleCorsMiddleware(http.HandlerFunc(handlers.GetProducts)))

	// mux.HandleFunc("/create-products", createProduct) //Old Routing
	// mux.Handle("POST /create-products", http.HandlerFunc(createProduct))
	mux.Handle("POST /create-products", global_router.HandleCorsMiddleware(http.HandlerFunc(handlers.CreateProduct)))

	fmt.Println("Server running on port http://localhost:3001")

	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
