package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	Articles = []Article{
		Article{Title: "Hello 1", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: getAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {

	R := mux.NewRouter()

	R.HandleFunc("/", homePage)
	R.HandleFunc("/all", getAllArticles)

	log.Fatal(http.ListenAndServe(":10000", R))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
