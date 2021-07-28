package main

import (
        "encoding/json"
        "net/http"
)

type Post struct {
     Title	string `json:"Title"`
     Author	string `json:"Author"` 
     Text	string `json:"Text"`
}

func main() {

    http.HandleFunc("/posts", postHandler)
    http.ListenAndServe(":8080", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
     posts := []Post {
          Post{"Post one", "John", "This is first post."},
          Post{"Post second", "Joh", "This is second post."},
          Post{"Post three", "Jo", "This is third post."},
         }

      en := json.NewEncoder(w)
      en.Encode(posts)

// and or json.NewEncoder(w).Encode(posts) check json newencoder func and encoder structure

}
     

