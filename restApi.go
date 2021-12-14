package main
import (
	"fmt" 
    "log"
    "net/http"
	"encoding/json"
)

type Article struct{
	Title string 'json:"Title"'
	Desc string 'json:"Desc"'
	Content string 'json:"Content"'

type Articles []Article

func allArticles(w http.ResponserWriter, r *http.Requesrt){
	articles := Articles{
		Articles{Title: "Test Title", Desc: "Test Description", Content: "Test Content"}
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
}
	 	                                                                                                                                                                                                                   
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home Page Hit")
}

func handleReuests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleReuests()
}


