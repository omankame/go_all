// this code works but problem is data gets overwrite in test.json

package main

import (
        "net/http"
        "html/template"
        "fmt"
        "encoding/json"
	"io/ioutil"
      )
var tpl *template.Template

type Employee struct {
     Name   string
     Gender string
     Age    string
}

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
     http.HandleFunc("/", index)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
     var emp Employee
     
     n := req.FormValue("name")
     g := req.FormValue("gender")
     a := req.FormValue("age")

     tpl.ExecuteTemplate(w , "index.gohtml", Employee{n, g, a})  

     emp = Employee{n, g, a}

     file , _ := json.MarshalIndent(emp, "", "")
     err := ioutil.WriteFile("test.json", file, 0644)
     if err != nil {
     fmt.Println("Error while creating and/or inserting data", err)
     return
     }          
}
