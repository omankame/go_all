package main

import (
        "fmt"
        "net/http"
        "html/template"
)

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
     http.HandleFunc("/", index)
     http.HandleFunc("/foo", foo)
     http.Handle("/favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}        

func index(w http.ResponseWriter, req *http.Request) {
     tpl.ExecuteTemplate(w, "index.gohtml" , nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     s := `Here is some text from foo`
     fmt.Fprint(w,s)
}     
 

