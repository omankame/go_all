  
package main

import (
	"database/sql"
	"net/http"
//	"log"
	_ "github.com/go-sql-driver/mysql"
        "fmt"

)

const hashCost = 8
var db *sql.DB

const msg1 = "Username and Password Added in the database sucessfully"
const msg2 = "Congrats Authentication done sucessfully"


func main() {
	// "Signin" and "Signup" are handler that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	// initialize our database connection
//	initDB()
	// start the server on port 8000
	http.ListenAndServe(":8080", nil)

}



func initDB() (db *sql.DB)  {
     db, err := sql.Open("mysql", "root:Golang@123@tcp(localhost:3306)/jio?charset=utf8")
     if err != nil {
        fmt.Println("err")
        panic("Could not continue as database connection not done")
    }


    return db
}



