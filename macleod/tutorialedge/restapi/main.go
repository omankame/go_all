package main

import (
        "fmt"
        "net/http"
        "encoding/json"
        
)
type Article struct {
     Title    string `json:"Title"`
     Desc     string `json:"Desc"`
     Content  string `json:"Content"`
}	

var Articles []Article

func main() {
    
 

            Articles = []Article{
                  Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
                  Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
            }
                 

    
     http.HandleFunc("/", homePage)
     http.HandleFunc("/articles", retrieveArticle)
     http.ListenAndServe(":8080", nil)


}

func homePage(w http.ResponseWriter, req *http.Request) {
     fmt.Fprint(w, "Welcome to Home Page")
     fmt.Println("Endpoint hit: homepage")

}

func retrieveArticle(w http.ResponseWriter, req *http.Request) {
      fmt.Println("Endpoint Hit: returnAllArticles")
      json.NewEncoder(w).Encode(Articles)

}
