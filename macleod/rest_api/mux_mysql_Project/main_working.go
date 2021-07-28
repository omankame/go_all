package main

import (
        "fmt"
        "encoding/json"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
        "net/http"
        "github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Post struct {
     Id     string `json: "id"`
     Title  string `json: "title"`
}

var posts []Post
func main() {

     r := mux.NewRouter()
     r.HandleFunc("/posts", getPosts).Methods("GET")
     http.ListenAndServe(":8080", r)
}



func getPosts(w http.ResponseWriter, req *http.Request) {
     db, err = sql.Open("mysql", "root:Redhat@123@tcp(localhost:3306)/test?charset=utf8")
     check(err)    
     defer db.Close()

     res, err := db.Query("select * from Post")
     check(err)
     defer res.Close()
     for res.Next() {
         var post Post
         err := res.Scan(&post.Id, &post.Title)
         check(err)
        posts = append(posts, post)
      }
         
     w.Header().Set("Content-Type", "application/json")
     json.NewEncoder(w).Encode(posts)

}     
     

func check(err error) {
     fmt.Println(err)
}
  

