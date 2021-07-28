package main

import (
        "fmt"
        "log"
        "strconv"
        "html/template"
        "net/http"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
)

const db_con_str = "root:Golang@123@tcp(localhost:33306)/jio?charset=utf8"

var tpl *template.Template
var db *sql.DB

type Employee struct {
     Id         int    `json="id"`
     FirstName  string `json="firstname"`
     LastName   string `json="lastname"`
     Email      string `json="email"`
    
}

func init() {

     tpl = template.Must(template.ParseGlob("templates/*"))
//     db  = dbConn()        
}

func dbConn() (db *sql.DB) {
     db, err := sql.Open("mysql", db_con_str)
     if err != nil {
        fmt.Println("Problem while connecting dataase")
        panic(err) 
        
     } 
     return db

}


func main() {
     http.HandleFunc("/", index)
     http.HandleFunc("/allemployee", allemp) 
     http.HandleFunc("/newemployee", newemp)
     http.HandleFunc("/deleteemployee", deleteemp)
     http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
     
     tpl.ExecuteTemplate(w, "index.gohtml", nil) 

}

func allemp(w http.ResponseWriter, req *http.Request) {
     db  = dbConn()
     defer db.Close()
     rows, err := db.Query("SELECT * FROM Employee")
     if err != nil {
        log.Println(err)
        return
     }

     emp := Employee{}   // instance of structure
     res := []Employee{}  // arrays of employee

     for rows.Next() {
         var id int
         var firstname, lastname, email string
         err = rows.Scan(&id, &firstname, &lastname, &email)
         if err != nil {
            log.Println(err)
            return
         }

         emp.Id = id
         emp.FirstName = firstname
         emp.LastName  = lastname
         emp.Email     = email

         res = append(res, emp)            
     }

     tpl.ExecuteTemplate(w, "allemployee.gohtml", res)     
}

func newemp(w http.ResponseWriter, req *http.Request) {

     db  = dbConn()
     defer db.Close()

     w.Header().Set("Content-Type", "text/html; charset=utf-8")   // for header info
     switch req.Method {
       case  http.MethodPost:
        req.ParseForm()        
        id := req.FormValue("eid") 
        firstn := req.FormValue("first")
        lastn := req.FormValue("last")
        eml := req.FormValue("email")
       
        eid, _  := strconv.Atoi(id)

//        emp := Employee{}
//        fmt.Println(eid, firstn, lastn, eml)
          if firstn == "" || lastn == "" || eml == "" {
             msg := "Enter nil value, not acceptable"
             tpl.ExecuteTemplate(w, "erroroccur.gohtml", msg)
             return
          }  
 
        insEmp, err :=  db.Prepare("INSERT INTO Employee(Id, FirstName, LastName, Email) VALUES(?,?,?,?)")
        if err != nil {
           db.Close()
           log.Println(err)
           tpl.ExecuteTemplate(w, "erroroccur.gohtml", err)
           return
        }

        res, err :=  insEmp.Exec(eid, firstn, lastn, eml)
        if err != nil {
              db.Close()
              log.Println(err)
              tpl.ExecuteTemplate(w, "erroroccur.gohtml", err)
              return
        }
        fmt.Println(res)
        tpl.ExecuteTemplate(w, "newemployee.gohtml", nil) 
        
     case http.MethodGet:
       tpl.ExecuteTemplate(w, "newemployee.gohtml", nil)
    }
       
}

       

func deleteemp(w http.ResponseWriter, req *http.Request) {

     db  = dbConn()
     defer db.Close()
     w.Header().Set("Content-Type", "text/html; charset=utf-8")   // for header info
     switch req.Method {
            case  http.MethodPost:
                      req.ParseForm()
                      id := req.FormValue("eid")
                      eid, _  := strconv.Atoi(id)
                       if eid == 0 {
                           msg := "You can not delete nil"
                           tpl.ExecuteTemplate(w, "erroroccur.gohtml", msg)
                           return
                       }
  
                       res, err :=  db.Exec("DELETE  FROM  Employee where Id = ?", eid)
                       if err != nil {
                       db.Close()
                       log.Println(err)
                       tpl.ExecuteTemplate(w, "erroroccur.gohtml", err)
                       return
                       }
                       fmt.Println(res)
                       
                       tpl.ExecuteTemplate(w, "deleteemployee.gohtml", nil)

            case  http.MethodGet:
                  tpl.ExecuteTemplate(w, "deleteemployee.gohtml", nil)
    }

}



 
