//https://gist.github.com/andreagrandi/97263aaf7f9344d3ffe6  very very important while post and put in json format


package main

import (
//        "fmt"
        "net/http"
        "github.com/gorilla/mux"
        "encoding/json"
//        "strconv"
//        "math/rand"
)

type Post struct {
     ID     string `json: "id"`
     Title  string `json: "title"`
     Body   string `json: "body"`
}

var posts []Post
func main() {

     posts = append(posts, Post{ID: "1", Title: "My first post", Body:"This is the content of my first post"})     
     posts = append(posts, Post{ID: "2", Title: "My second post", Body:"This is the content of my second post"})     

     r := mux.NewRouter()
     r.HandleFunc("/posts", getPosts).Methods("GET")
     r.HandleFunc("/posts/{id}", getPost).Methods("GET")
     r.HandleFunc("/posts", createPost).Methods("POST")
     r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
     r.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
     http.ListenAndServe(":8080", r)
}

func getPosts(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     json.NewEncoder(w).Encode(posts) 

}

func getPost(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     param := mux.Vars(req)
     for _, item := range posts {
        if item.ID == param["id"] {
           json.NewEncoder(w).Encode(item)
           return
         }
      }
    json.NewEncoder(w).Encode(&Post{})
}     

func createPost(w http.ResponseWriter, req *http.Request) {

     w.Header().Set("Content-Type", "application/json")
     var post Post
     decoder := json.NewDecoder(req.Body)
     _ = decoder.Decode(&post)
//     post.ID = strconv.Itoa(rand.Intn(100))
     posts = append(posts, post)
//     json.NewEncoder(w).Encode(&post)
      json.NewEncoder(w).Encode(posts)
}

func deletePost(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
//     var post Post
//     decoder := json.NewDecoder(req.Body)
//     _ = decoder.Decode(&post)
       param := mux.Vars(req)
       for ind, item := range posts {
          if item.ID == param["id"] {
             posts = append( posts[:ind], posts[ind+1:]...)
             break
         }
      }      
     json.NewEncoder(w).Encode(posts)
}

func updatePost(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     param := mux.Vars(req)
       for ind, item := range posts {
          if item.ID == param["id"] {
             posts = append( posts[:ind], posts[ind+1:]...)

            var post Post
            decoder := json.NewDecoder(req.Body)
            _ =  decoder.Decode(&post)
            post.ID = param["id"]
//            post.Title = param["title"]
            posts = append(posts, post)
            json.NewEncoder(w).Encode(&post)
            return
         }
      }
     json.NewEncoder(w).Encode(&Post{})
}
   





