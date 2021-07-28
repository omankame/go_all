package main

import (
       "io"
       "net/http"
       "html/template"
       "log"
)

var tpl *template.Template

func main() {

     http.HandleFunc("/", foo)
     http.HandleFunc("/dog/", dog)
     http.HandleFunc("/dog.jpg", dg)
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "text/html; charset=UTF-8")
     io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
        tpl, err :=  template.ParseFiles("dog.gohtml")
        if err != nil {
        log.Fatalln(err)
        }

      tpl.ExecuteTemplate(w , "dog.gohtml", nil)
}

func dg(w http.ResponseWriter, req *http.Request) {
     http.ServeFile(w , req, "dog.jpg")
}
