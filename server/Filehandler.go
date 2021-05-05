package main

import (
	"github.com/gorilla/mux"
	"image/png"
	"net/http"
	"os"
)

func sendPhotosToClient(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets param
	path := "./img/"+params["name"]

	file, err := os.Open(path)
	defer	file.Close()
	if err != nil {
		http.NotFound(w,r)
		return
	}

	img, err := png.Decode(file)
	if err != nil {
		http.NotFound(w,r)
		return
	}

	png.Encode(w, img) // Write to the ResponseWriter
}