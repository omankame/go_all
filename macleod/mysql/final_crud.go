package main

import (
        "html/template"
        "net/http"
        "log"
        "fmt"
        "database/sql"
       _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
     Empid       int
     First_name  string
     Last_name   string
}

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*"))
}

func dbConn() (db *sql.DB) {
     db, err := sql.Open("mysql", "root:Redhat@123@tcp(localhost:3306)/test?charset=utf8")
     if err != nil {
        log.Fatalln(err)
     }
     return db
}

func main() {
     http.HandleFunc("/", index)
     http.Handle("/favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

     db := dbConn() 
     
     selDb, err := db.Query("select * from employees")
     if err != nil {
        fmt.Println(err)
     }
     emp := Employee{}
     res := []Employee{} 

     for selDb.Next() {
         var empid   int
         var first_name, last_name string
         err := selDb.Scan(&empid, &first_name, &last_name)
         if err != nil {
            log.Println(err)
         }
       emp.Empid = empid
       emp.First_name = first_name
       emp.Last_name = last_name

       res = append(res,emp)
     }
  
     tpl.ExecuteTemplate(w, "index.gohtml", res) 

}
