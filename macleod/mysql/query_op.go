package main

import ( 
        "io"
        "net/http"
        "database/sql"
        "fmt"
        _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

//type Employee struct {
//     empid	int
//     first_name string
//     last_name  string
//} 

func main() {
     db, err = sql.Open("mysql", "root:Redhat@123@tcp(localhost:3306)/test?charset=utf8")
     check(err)
     defer db.Close()

     err = db.Ping()
     check(err)
     
     http.HandleFunc("/", index)
     http.HandleFunc("/amigos", amigos)
     http.HandleFunc("/create", create)
     http.HandleFunc("/insert", insert)
     http.HandleFunc("/read", read)
     http.HandleFunc("/update", update)
     http.HandleFunc("/delete", del)
     http.HandleFunc("/drop", drop)
     http.Handle("/favicon.ico", http.NotFoundHandler())
     err := http.ListenAndServe(":8080", nil)
     check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
     io.WriteString(w, "AT INDEX")

}

func amigos(w http.ResponseWriter, req *http.Request) {
     rows, err := db.Query(`Select first_name from employees;`)
     check(err)
     var s,  fname  string
     s = "RETRIEVED RECORDS:\n"

     for rows.Next() {
          err = rows.Scan( &fname)
          check(err)
          s +=  fname +  "\n"
          }
     fmt.Fprint(w, s)     
}

func create(w http.ResponseWriter, req *http.Request) {
      query := `CREATE table customer (name varchar(20));`  
      res, err := db.Query(query) 
      if err != nil {
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
      }
      defer res.Close()
    
     io.WriteString(w, "Table has been created") 
}



func insert(w http.ResponseWriter, req *http.Request) {
     query := `Insert into customer value ( 'ONKAR');`
     stmt, err := db.Prepare(query)
     check(err)
     defer stmt.Close()

     res, err := stmt.Exec()
     check(err)
     n, err := res.RowsAffected()
     fmt.Fprintln(w, "Value added into customer", n)          
 
}

func read(w http.ResponseWriter, req *http.Request) {

}

func update(w http.ResponseWriter, req *http.Request) {

}

func del(w http.ResponseWriter, req *http.Request) {

}

func drop(w http.ResponseWriter, req *http.Request) {

     query := `DROP TABLE customer;`
     stmt, err := db.Prepare(query) 
     check(err)
     defer stmt.Close()

     res, err := stmt.Exec()
     check(err)

     fmt.Fprint(w, "Table Dropped", res)
}    


func check(err error) {
     if err != nil {
     fmt.Println(err)
     }
}

// https://stackoverflow.com/questions/16280176/go-panic-runtime-error-invalid-memory-address-or-nil-pointer-dereference  is solution of null
// null pointer exception.


