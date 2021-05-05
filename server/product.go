package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products .....
type Products struct {
	IdProduct     int  `json:"id"`
	Name   string  `json:"name"`
	Price  float64  `json:"price"`
	Picture string `json:"picture"`
	Categories_id int `json:"categories_id"`
}

// Get all Product
func listProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var returnProducts []Products
	categoryGet, ok1 := r.URL.Query()["category"]
	minpriceGet, ok2 := r.URL.Query()["minPrice"]
	maxPriceGet, ok3 := r.URL.Query()["maxPrice"]
	nameGet, ok4 := r.URL.Query()["name"]

	sqlWhere := "where 1 = 1"

	if ok1 && len(categoryGet[0]) >= 1 {
		catId, _ := strconv.Atoi(categoryGet[0])
		sqlWhere += " AND categories_id = " + strconv.Itoa(catId)
	}
	if ok2 && len(minpriceGet[0]) >= 1 {
		minPrice, _ := strconv.Atoi(minpriceGet[0])
		sqlWhere += " AND price > " + strconv.Itoa(minPrice)
	}
	if ok3 && len(maxPriceGet[0]) >= 1 {
		maxPrice, _ := strconv.Atoi(maxPriceGet[0])
		sqlWhere += " AND price < " + strconv.Itoa(maxPrice)
	}
	if ok4 && len(nameGet[0]) >= 1 {
		sqlWhere += " AND name like '%" + nameGet[0] + "%'"
	}

	result, err := db.Query("SELECT * from product " + sqlWhere)
	defer result.Close()

	for result.Next() {
		var product Products
		result.Scan(&product.IdProduct, &product.Name, &product.Price, &product.Picture, &product.Categories_id)
		returnProducts = append(returnProducts, product)
	}
	err = json.NewEncoder(w).Encode(returnProducts)
	if err != nil {
		return 
	}
}

// Get single Product
func listOneProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Gets params
	id, _ := strconv.Atoi(params["id"])
	w.Header().Set("Content-Type", "application/json")
	var retprod []Products

	result, _ := db.Query("SELECT * from product where idProduct = ?", id)
	defer result.Close()

	for result.Next() {
		var product Products
		result.Scan(&product.IdProduct, &product.Name, &product.Price, &product.Picture, &product.Categories_id)
		retprod = append(retprod, product)
	}
	json.NewEncoder(w).Encode(retprod)

}

// Add new Product
func addProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p Products
	json.NewDecoder(r.Body).Decode(&p)

	db.Exec("INSERT INTO `shop`.`product` (`name`, `price`, `picture`, `categories_id`) VALUES (?, ?, ?, );", p.Name, p.Price, p.Picture, p.Categories_id)
	json.NewEncoder(w).Encode(p)
}

// Update Product
func changeProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Products
	 json.NewDecoder(r.Body).Decode(&product)

	res,_ := db.Exec("UPDATE `shop`.`product` SET `name` = ?, `price` = ?, `picture` = ?, `categories_id` = ? WHERE (`idProduct` = ?);",
		product.Name, product.Price, product.Picture, product.Categories_id, product.IdProduct)

	json.NewEncoder(w).Encode(res)


}

// Delete Product
func removeProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	res,_ := db.Exec("DELETE FROM `shop`.`product` WHERE (`idProduct` = ?);) VALUES (?, ?, ?, );", id)

	json.NewEncoder(w).Encode(res)
	}


