package handlers

import (
	"ecommerce/database"
	"ecommerce/global_router"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	global_router.HandleCors(w)
	global_router.HandlePreflightReq(w, r)

	util.SendData(w, database.ProductList, 200)
}
