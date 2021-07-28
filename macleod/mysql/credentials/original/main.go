package main

import (
        "net/http"
        "fmt"
        "database/sql"
      _ "github.com/go-sql-driver/mysql"  
        "html/template"
)

type Credentials struct {
     Username  string `json:"username", db:"username"` 
     Password  string `json:"password", db:"password"`

}
var (
      Db *sql.DB
      tpl *template.Template
)
const msg1 = "Username and Password Added in the database sucessfully"
const msg2 = "Congrats Authentication done sucessfully"

func init() {
          tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
     
     http.HandleFunc("/signup", signup)
     http.HandleFunc("/signin", signin)
     http.ListenAndServe(":8080", nil)

}

func initDB() (Db *sql.DB)  {
     Db, err := sql.Open("mysql", "root:Golang@123@tcp(localhost:3306)/jio?charset=utf8")
     if err != nil {
        fmt.Println("err")
        panic("Could not continue as database connection not done")
    }

    
    return Db   
}
     


