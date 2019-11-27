package handler

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/saiprasanna/dbinsert"
	"strconv"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

var db *sql.DB

var err error

var photocolumns = map[string]func(string) (interface{}, error){
	"id":      atoi,
	"albumId": atoi,
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var albums []dbinsert.Album

	db, err = sql.Open("mysql", "root:infoblox@tcp(127.0.0.1:3306)/typicode")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	result, err := db.Query("SELECT id, userId, title from album")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var album dbinsert.Album
		err := result.Scan(&album.Id, &album.UserId, &album.Title)
		if err != nil {
			panic(err.Error())
		}
		albums = append(albums, album)
	}
	json.NewEncoder(w).Encode(albums)
}

func atoi(s string) (interface{}, error) {
	return strconv.Atoi(s)
}

func GetPhotos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var photos []dbinsert.Photo

	db, err = sql.Open("mysql", "root:infoblox@tcp(127.0.0.1:3306)/typicode")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	where := ""
	params := []interface{}{}
	for k, v := range r.URL.Query() {
		if convert, ok := photocolumns[k]; ok {
			param, err := convert(v[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			params = append(params, param)
			where += k + " = ? AND "
		}
	}
	if len(where) > 0 {
		// prefix the string with WHERE and remove the last " AND "
		where = " WHERE " + where[:len(where)-len(" AND ")]
	}

	result, err := db.Query("SELECT id, albumId, Title,thumbnailUrl,url from photo"+where, params...)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var photo dbinsert.Photo
		err := result.Scan(&photo.Id, &photo.AlbumId, &photo.Title, &photo.Url, &photo.Thumbnailurl)
		if err != nil {
			panic(err.Error())
		}
		photos = append(photos, photo)
	}
	json.NewEncoder(w).Encode(photos)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	db, err = sql.Open("mysql", "root:infoblox@tcp(127.0.0.1:3306)/typicode")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	result, err := db.Query("SELECT id, userId, title from album where id=?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var album dbinsert.Album

	for result.Next() {
		err := result.Scan(&album.Id, &album.UserId, &album.Title)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(album)
}
