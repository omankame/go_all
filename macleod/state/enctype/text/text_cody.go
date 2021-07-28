package main

import (
//        "fmt"
        "html/template"
        "net/http"
)

var tpl *template.Template

type person struct {
     FirstName	string
     LastName	string
     Subscribe	bool	
}
func init() {
    tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
     http.HandleFunc("/", foo)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     bs := make([]byte, req.ContentLength)
     req.Body.Read(bs)
     body := string(bs)

    tpl.ExecuteTemplate(w, "index.gohtml", body) 
}
