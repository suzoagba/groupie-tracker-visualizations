package main

import (
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/artists"
	"groupie-tracker/handler"
	"log"
	"net/http"
)

func main() {
	// Parse the API JSON's
	api.GetFromApi()

	fmt.Println("- Linking the relations API and artists API")
	artists.LinkArtists()

	// Start the server
	http.HandleFunc("/", handler.Handler)
	http.HandleFunc("/favicon.ico", handler.FaviconHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js/"))))

	fmt.Println("\nFinished!\nGo to: http://localhost:8080")
	handler.Open("http://localhost:8080/")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
