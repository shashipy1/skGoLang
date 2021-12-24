package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title         string `json:"Title"`
	NAME          string `json:"NAME"`
	Myself        string `json:"Myself"`
	Desription    string `"json:Desription"`
	Cotactno      string `json:"Cotactno"`
	Facebooklink  string `json:"Facebooklink"`
	Instagramlink string `json:"Instagramlink"`
	Linkdinlink   string `json:"Linkdinlink"`
	GitHubLink    string `json:"GitHubLink"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	myself := Articles{
		Article{Title: "About My Self", NAME: "SHASHI KANT KUMAR", Myself: "It's my pleasure to introduce myself. I am pursuing B. Tech in CSE from Radharaman Institute of Technology & Science, Bhopal. As far as concerned my family, I belong to meddle class family. My father is a compounder and my mothor is homemaker. My younger brother and sister are preparing general. I am shashi kant kumar. I belong to Bihar.I have teached class 5 to 10 six month in private school.I am a good in programming language c, c++, python and Golang and very much interested in HTML, CSS and JS.My strength is self confidence, positive attitude and hard worker. My weakness is I am easily believe everyone My hobbies are playing chess, read Scripture book and listening prabhupada lecture and spiritual songs.", Desription: "Hey Shashi kant this side a purshing B.Tech 3rd year in Computer Science And Engimeering. I have six month experince of web devlopment through Python and Django, also know about Golang React Java and so many skils. I have done six month SEO work on BRNEWS a quick learner slove daily practic problem on hackerrankwith 3 stat", Cotactno: "700495149 9135004608", Facebooklink: "https://www.facebook.com/profile.php?id=100015181350174", Instagramlink: "https://www.instagram.com/ranjha__217/", Linkdinlink: "https://www.linkedin.com/in/s-k-ranjha-08b4a0226/", GitHubLink: "https://github.com/shashipy1"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(myself)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Introducing my self: It's my pleasure to introduce myself. I am shashi kant kumar. I belong to Bihar. I am pursuing B. Tech in CSE from RITS, Bhopal. As far as concerned my family, I belong to meddle class family. My father is a compounder and my mothor is homemaker. My younger brother and sister are preparing general. I have teached class 5 to 10 six month in private school. I am a good in programming language c, c++, python and Golang and very much interested in HTML, CSS and JS. My strength is self confidence, positive attitude and hard worker. My weakness is I am easily believe everyone. My hobbies are playing chess, read Scripture book and listening prabhupada lecture and spiritual songs.Thank you")
}

func handleReuests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/myself", allArticles)

	log.Fatal(http.ListenAndServe(":7009", nil))
}

func main() {
	handleReuests()
}
