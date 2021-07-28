package main

import (
        "fmt"
        "net/http"
        "encoding/json"
        "github.com/gorilla/mux"        
        "io/ioutil"
)
type Article struct {
     Id       string `json:"Id"`
     Title    string `json:"Title"`
     Desc     string `json:"Desc"`
     Content  string `json:"Content"`
}	

var Articles []Article

func main() {
    
 

            Articles = []Article{
                  Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
                  Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
            }
                 
     myrouter := mux.NewRouter().StrictSlash(true)
   
     myrouter.HandleFunc("/", homePage)
     myrouter.HandleFunc("/all", retrieveArticle)
     myrouter.HandleFunc("/article/{id}", returnSingleArticle)
     myrouter.HandleFunc("/article", createArticle).Methods("POST")
     http.ListenAndServe(":8080", myrouter)


}

func homePage(w http.ResponseWriter, req *http.Request) {
     fmt.Fprint(w, "Welcome to Home Page")
     fmt.Println("Endpoint hit: homepage")

}

func retrieveArticle(w http.ResponseWriter, req *http.Request) {
      fmt.Println("Endpoint Hit: returnAllArticles")
      json.NewEncoder(w).Encode(Articles)

}

func returnSingleArticle(w http.ResponseWriter, req *http.Request) {

     vars := mux.Vars(req)
     key  := vars["id"]
//     fmt.Fprint(w, "Key: " + key)

     for _, artstr := range Articles {
         if key == artstr.Id {
            json.NewEncoder(w).Encode(artstr)
         }
     }
}
 
func createArticle(w http.ResponseWriter, req *http.Request) {
     
     reqBody, _ := ioutil.ReadAll(req.Body)
     var article  Article
     json.Unmarshal(reqBody, &article)
    
     Articles = append(Articles, article)   
  
     json.NewEncoder(w).Encode(Articles)
//     fmt.Fprintf(w, "%+v", string(reqBody))     

}
