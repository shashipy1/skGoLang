package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
	Contact string `json:"Contact"`
	Name    string `json:"Name"` 
	add     string `json: "add"`

} 

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Test Content", Contact: "9135004608", Name: "SHASHI KANT KUMAR", add: "Vaishali Nagar77, Bhopal(MP)"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w, "Home Page Hit")
}

func handleReuests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)

	log.Fatal(http.ListenAndServe(":8988", nil))
}

func main() {
	handleReuests()
}
