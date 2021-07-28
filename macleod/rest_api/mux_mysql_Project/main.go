// important post https://medium.com/@hugo.bjarred/rest-api-with-golang-mux-mysql-c5915347fa5b
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

    db, err = sql.Open("mysql", "root:Redhat@123@tcp(localhost:3306)/test?charset=utf8")
    check(err)
    defer db.Close()


     r := mux.NewRouter()
     r.HandleFunc("/posts", getPosts).Methods("GET")
     r.HandleFunc("/posts/{id}", getPost).Methods("GET")
     r.HandleFunc("/posts", createPost).Methods("POST")
     r.HandleFunc("/", RestAPI)
     http.ListenAndServe(":8080", r)
}
func RestAPI(w http.ResponseWriter, req *http.Request) {
     str := fmt.Sprint(`<h1>Welcome To RESTAPI + MYSQL Evaluation </h1>` ) 
     fmt.Fprint(w, str)
}

func getPosts(w http.ResponseWriter, req *http.Request) {
     
     posts := dbLogic()
     w.Header().Set("Content-Type", "application/json")
     json.NewEncoder(w).Encode(posts)

} 

func getPost(w http.ResponseWriter, req *http.Request) {
     posts := dbLogic()
     w.Header().Set("Content-Type", "application/json")
     param := mux.Vars(req)
     for _, item := range posts {
          if param["id"] ==  item.Id {
           json.NewEncoder(w).Encode(item)
           return
          }
     }
     json.NewEncoder(w).Encode(&Post{})  

}    

func createPost(w http.ResponseWriter, req *http.Request) {
     var post Post
     w.Header().Set("Content-Type", "application/json")
     decoder := json.NewDecoder(req.Body)
     _ = decoder.Decode(&post)
     query := `insert into Post (id, title) values (?);`
     stmt, err := db.Prepare(query)
     check(err)
     defer stmt.Close() 
     res, err := stmt.Exec()
     check(err)

     http.Redirect(w , req, "/" , http.StatusMovedPermanently)
     
}
     
func dbLogic() []Post {
     var posts []Post
     res, err := db.Query("select * from Post")
     check(err)
     defer res.Close()
     for res.Next() {
         var post Post
         err := res.Scan(&post.Id, &post.Title)
         check(err)
        posts = append(posts, post)
      }
  return posts
}

func check(err error) {
     fmt.Println(err)
}
  

