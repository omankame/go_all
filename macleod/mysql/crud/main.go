// https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html
package main

import (
        "net/http"
        "html/template"
        "fmt"
        "database/sql"
        _ "github.com/go-sql-driver/mysql" 
        "strconv"
//        "io"
)

type Datainfo struct {
     Id 	int
     Name	string
     City	string
}

var tpl *template.Template
var err error
var db *sql.Conn

func init() {
     tpl = template.Must(template.ParseGlob("form/*"))
}

func main() {
     http.HandleFunc("/", Index)
     http.HandleFunc("/show", Show)
     http.HandleFunc("/delete", Delete)
     http.HandleFunc("/edit", Edit)
     http.HandleFunc("/update", Update)  
     http.HandleFunc("/insert", Insert)
     http.HandleFunc("/new", New)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func dbConn() (db *sql.DB) {
     db, err = sql.Open("mysql", "root:Redhat@123@tcp(localhost:3306)/test?charset=utf8")
     check (err)
     return db
}

func check(err error) {
     if err != nil {
     fmt.Println(err)
    }
}

func New(w http.ResponseWriter, r *http.Request) {
    tpl.ExecuteTemplate(w, "New", nil)
}

func Show(w http.ResponseWriter, req *http.Request) {
     num, err := strconv.Atoi(req.URL.Query().Get("id"))
     check(err)
//     fmt.Fprint(w, id)
 
      conn := dbConn()
      defer conn.Close()
      data := Datainfo{}
//      final := []Datainfo{}
      res, err := conn.Query("select * from datainfo where id = ?", num)
      check(err)
      defer res.Close()
    
      for res.Next() {
          var id int
          var name, city string
          res.Scan(&id, &name, &city)
          data.Id = id
          data.Name = name
          data.City = city

  //        final = append(final, data)
          }
     tpl.ExecuteTemplate(w, "Show", data)
     fmt.Println("Executed template")
}

func Insert(w http.ResponseWriter, req *http.Request) {
     conn := dbConn()
     

     w.Header().Set("Content-Type", "text/html; charset=utf-8")
     if req.Method == http.MethodPost {
         name := req.FormValue("name")
         city := req.FormValue("city")
         insForm, err := conn.Prepare("INSERT INTO datainfo(name, city) VALUES(?,?)")
         check(err)
         insForm.Exec(name, city)
         fmt.Println("INSERT: Name: " + name + " | City: " + city)
       }
     defer conn.Close()
     http.Redirect(w, req, "/", 301)  
     
/*     io.WriteString(w, `
        <form method="POST">
         <input type="text" name="q">
         <input type="submit">
        </form>
        <br>`+v)*/
}

func Update(w http.ResponseWriter, req *http.Request) {
     conn := dbConn()
     w.Header().Set("Content-Type", "text/html; charset=utf-8")
     if req.Method == http.MethodPost {
     name := req.FormValue("name")
     city := req.FormValue("city")
     id := req.FormValue("uid")
 
     insForm, err := conn.Prepare("Update datainfo set name=?, city=? where id=?")
     check(err)
     insForm.Exec(name, city, id)
     fmt.Println("UPDATE: Name: " + name + " | City: " + city)
     }
     defer conn.Close()
     http.Redirect(w, req, "/", 301)
} 
 
func Edit(w http.ResponseWriter, req *http.Request) {
     num, err := strconv.Atoi(req.URL.Query().Get("id"))
     check(err)
     conn := dbConn()
     defer conn.Close()
     data := Datainfo{}
     
     res, err :=  conn.Query("select * from datainfo where id = ?", num)
     check(err) 
  
     for res.Next() {
         var id int
         var name, city string
         res.Scan(&id, &name, &city)
         data.Id =  id
         data.Name = name
         data.City = city
         }
       tpl.ExecuteTemplate(w, "Edit", data)
}
func Delete(w http.ResponseWriter, req *http.Request) {
     num, err := strconv.Atoi(req.URL.Query().Get("id"))
     check(err)
  
     conn := dbConn()
     defer conn.Close()
     _, err = conn.Exec("delete from datainfo where id = ?", num)
     check(err)

      http.Redirect(w , req, "/", http.StatusMovedPermanently)
      fmt.Println("Executed template")
}

     

func Index(w http.ResponseWriter, req *http.Request) {
 
//    dinf1 := Datainfo {
//            Id:	1,
//            Name: "Onkar",
//            City: "Thane",
//           }
//    dinf2 := Datainfo {
//            Id: 2,
//            Name: "Nikhil",
//            City: "Alibag",
//           }

//    res := []Datainfo{}  or res := []Datainfo{}
//    res = append(res, dinf1)
//    res = append(res, dinf2)
   
      conn := dbConn() 
      defer conn.Close()
      data := Datainfo{}
      final := []Datainfo{} 
      res, err := conn.Query("select * from datainfo") 
      check(err)
      defer res.Close()
    
      for res.Next() {
          var id int
          var name, city string
          res.Scan(&id, &name, &city)
          data.Id = id
          data.Name = name
          data.City = city

          final = append(final, data)
          }     
     tpl.ExecuteTemplate(w, "Index", final)
     fmt.Println("Executed template")
}
        
