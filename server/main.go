package main

import (
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

	mux.HandleFunc("/hello", helloHandle)

	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on port :3001")

	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
