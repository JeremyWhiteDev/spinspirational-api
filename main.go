package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Album represents data about a record Album.
type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/album", getAlbums)
	mux.HandleFunc("/album/{id}", getAlbumById)
	mux.HandleFunc("POST /album", postAlbum)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	albumJson, err := json.Marshal(albums)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(albumJson))
}

func getAlbumById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	for _, a := range albums {
		if a.ID == id {
			albumJson, err := json.Marshal(a)
			if err != nil {
				panic(err)
			}
			fmt.Fprint(w, string(albumJson))
			return
		}
	}
	http.Error(w, "Resource not found", http.StatusNotFound)
}

func postAlbum(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close() 
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var newAlbum Album

	err = json.Unmarshal(body, &newAlbum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	albums = append(albums, newAlbum)

	newAlbumJson, err := json.Marshal(newAlbum)
	if err != nil {
		http.Error(w, "Error parsing JSON body", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(newAlbumJson))
}