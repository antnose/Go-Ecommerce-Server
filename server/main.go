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
	ID          int
	Title       string
	Description string
	Price       float64
	ImgURL      string
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Please give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandle)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running on port :3001")

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
