package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello I am Ibrahim. I Will be software engineer")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Please give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Please make an POST method", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandle)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/create-products", createProduct)

	fmt.Println("Server running on port http://localhost:3001")

	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	pd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is orange. And that's why it's orange",
		Price:       230,
		ImgURL:      "https://upload.wikimedia.org/wikipedia/commons/thumb/4/43/Ambersweet_oranges.jpg/250px-Ambersweet_oranges.jpg",
	}

	pd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is apple",
		Price:       340,
		ImgURL:      "https://hips.hearstapps.com/hmg-prod/images/apples-at-farmers-market-royalty-free-image-1627321463.jpg?crop=1xw:0.94466xh;center,top&resize=1200:*",
	}

	pd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is banana",
		Price:       40,
		ImgURL:      "https://www.allrecipes.com/thmb/jYmw-0Vijg1E_OuG2yGjEAcdQg4=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-4x3-d8f0871e12214350be7ae5575eea4eed.jpg",
	}

	productList = append(productList, pd1)
	productList = append(productList, pd2)
	productList = append(productList, pd3)

}
