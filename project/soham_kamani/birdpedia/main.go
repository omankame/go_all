package main

import (
        "fmt"
        "net/http"
        "github.com/gorilla/mux"
        _ "github.com/lib/pq"
        "database/sql"         
)

const (
hostname = "localhost"
host_port = 5432
username = "postgres"
password = "postgres"
databasename = "bird_encyclopedia"
)




func main() {
      connStr := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", hostname, host_port, username, password, databasename)
     db, err := sql.Open("postgres", connStr)
     if err != nil {
        panic(err)
     }

     db.Ping()

     InitStore(&dbStore{
                    db: db,}) 



     r := newRouter()

//     r.HandleFunc("/", handler).Methods("GET")

     http.ListenAndServe(":8080", r)

}

func newRouter() *mux.Router {
     r := mux.NewRouter()
     r.HandleFunc("/hello", handler).Methods("GET")

     staticFileDir := http.Dir("./assets")
     staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))
     r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
     r.HandleFunc("/bird", getBirdHandler).Methods("GET")
     r.HandleFunc("/bird", createBirdHandler).Methods("POST")
     return r
}

func handler(w http.ResponseWriter, req *http.Request) {
     fmt.Fprintf(w, "Hello World!") 
}


