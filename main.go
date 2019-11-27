package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/saiprasanna/dbinsert"
	"github.com/saiprasanna/handler"
)


func main() {

	url := "https://jsonplaceholder.typicode.com/albums"
	dbinsert.Insertalbum(url)
	url1 := "https://jsonplaceholder.typicode.com/photos"
	dbinsert.Insertphotos(url1)

	router := mux.NewRouter()
	router.HandleFunc("/albums", handler.GetAlbums).Methods("GET")
	router.HandleFunc("/albums/{id}", handler.GetAlbum).Methods("GET")
	router.HandleFunc("/photos", handler.GetPhotos).Methods("GET")

	http.ListenAndServe("10.196.105.125:8000", router)
}


