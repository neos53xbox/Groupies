package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	API "groupetracker/api"
	Template "groupetracker/template"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/artist/static/", http.StripPrefix("/artist/static/", fs))

	http.HandleFunc("/", Template.HandleTemplateMainPage)
	for id := range API.ArtistsList() {
		http.HandleFunc("/artist/"+strconv.Itoa(id+1), Template.HandleTemplateArtistPage)
	}

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
