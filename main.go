package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/albums", returnAllAlbums)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func returnAllAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllAlbums")
	json.NewEncoder(w).Encode(albums)
}

func main() {
	fmt.Println("Running server on: http://localhost:8081")
	handleRequests()
}
