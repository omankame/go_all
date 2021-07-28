package main

import (
        "net/http"
        "database/sql"
        "fmt"
        "io"
      _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
     db, err = sql.Open("mysql", "root:Redhat@123@tcp(localhost:3306)/test?charset=utf8")
     check (err)
     
     defer db.Close()
     err = db.Ping()
     if err != nil {
     panic(err.Error())
     }     

    var version string
    err = db.QueryRow("select version()").Scan(&version)
    check(err)

   fmt.Println(version)

   http.HandleFunc("/", index)
   http.Handle("favicon.ico", http.NotFoundHandler())
   http.ListenAndServe(":8080", nil)
}

func check(err error) {
     if err != nil {
     fmt.Println(err)
    }
}

func index(w http.ResponseWriter, req *http.Request) {
     _, err = io.WriteString(w, "Successfully Connected.") 
}
