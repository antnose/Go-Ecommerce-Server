package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	productId, err := strconv.Atoi(strId)

	if err != nil {
		http.Error(w, "Please give a valid product id", 400)
		return
	}

	for _, product := range database.ProductList {

		if productId == product.ID {
			util.SendData(w, product, 200)
			return
		}

	}

	util.SendData(w, "Product not Found", 404)
}
