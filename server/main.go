package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welche to the API")
}

func handleRequests() {

	r := mux.NewRouter().StrictSlash(true)
	godotenv.Load()

	db, _ = sql.Open("mysql", os.Getenv("databaseString"))
	defer db.Close()



	r.HandleFunc("/", homePage)
	r.HandleFunc("/login", login)
	r.HandleFunc("/img/{name}", sendPhotosToClient)
	r.HandleFunc("/product", listProducts).Methods("GET")
	r.HandleFunc("/product/{id}", listOneProducts).Methods("GET")
	r.HandleFunc("/product", addProducts).Methods("POST")
	r.HandleFunc("/product/{id}", changeProducts).Methods("PUT")
	r.HandleFunc("/product/{id}", removeProducts).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":10000", r))
}

func main() {
	fmt.Println("Rest API for Mini EC System")

	handleRequests()
}