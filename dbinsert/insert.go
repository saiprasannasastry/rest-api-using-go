package dbinsert

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

var db *sql.DB

var err error

type Album struct {
        Title  string `json:"title"`
        UserId int    `json:"userId"`
        Id     int    `json:"id"`
}

type Photo struct {
        Title        string `json:"title"`
        AlbumId      int    `json:"albumId"`
        Id           int    `json:"id"`
        Thumbnailurl string `json:"thumbnailUrl"`
        Url          string `json:"url"`
}


func Insertalbum(url string) {

	db, err = sql.Open("mysql", "root:infoblox@tcp(127.0.0.1:3306)/typicode")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var albumObject []Album

	json.Unmarshal([]byte(body), &albumObject)

	for _, value := range albumObject {
		insert, err := db.Query("INSERT IGNORE INTO album(id,UserId,Title) VALUES(?,?,?)", value.Id, value.UserId, value.Title)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
	}

}

func Insertphotos(url string) {

	db, err = sql.Open("mysql", "root:infoblox@tcp(127.0.0.1:3306)/typicode")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	req, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var photoObject []Photo

	json.Unmarshal([]byte(body), &photoObject)

	for _, value := range photoObject {
		_, err := db.Exec("INSERT IGNORE INTO photo(id,albumId,thumbnailUrl,Title,url) VALUES(?,?,?,?,?)", value.Id, value.AlbumId, value.Thumbnailurl, value.Title, value.Url)
		if err != nil {
			panic(err.Error())
		}

	}

}
