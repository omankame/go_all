package main

import ( 
//        "io"
//        "net/http"
        "database/sql"
        "fmt"
        _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type Employee struct {
     empid	int
     first_name string
     last_name  string
} 

func main() {
     db, err = sql.Open("mysql", "root:Redhat@123@tcp(localhost:3306)/test?charset=utf8")
     check(err)
     defer db.Close()

     err = db.Ping()
     check(err)
  
     res, err := db.Query("select * from employees")
     check(err)
     defer res.Close()
     
     for res.Next() {
       var emp Employee 
       err := res.Scan(&emp.empid, &emp.first_name, &emp.last_name)
       if err != nil {
         fmt.Println(err)
       }

      fmt.Printf("%v\n", emp)
    } 

         
}

func check(err error) {
     if err != nil {
     fmt.Println(err)
     }
}
