package main

import (
        "fmt"
        "net/http"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
        "html/template"
)



var db *sql.DB
var tpl *template.Template
const (
      hashCost = 8
      msg1 = "Username and Password Added in the database sucessfully"
      msg2 = "Congrats Authentication done sucessfully"
)

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

     http.HandleFunc("/signup", signup)
     http.HandleFunc("/login", login)
     http.ListenAndServe(":8080", nil)

}


func initDB() (db *sql.DB) {
     db, err := sql.Open("mysql", "root:Golang@123@tcp(localhost:3306)/jio?charset=utf8")  
     if err != nil {
        fmt.Println("Serious Problem, Could not connect to database")
        return 
     }   

    return db
}
