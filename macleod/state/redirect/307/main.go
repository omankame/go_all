package main

import (
        "fmt"
        "net/http"
        "html/template"
)
var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
     http.HandleFunc("/", foo)
     http.HandleFunc("/bar", foobar)
     http.HandleFunc("/barred", zoo)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)   
}

func foo(w http.ResponseWriter, req *http.Request) {
     fmt.Println("Requets Method in foo is:", req.Method, "\n")
}

func foobar(w http.ResponseWriter, req *http.Request) {
     fmt.Println("Requets Method in foo is:", req.Method, "\n")
     http.Redirect(w , req, "/", http.StatusTemporaryRedirect)     
}

func zoo(w http.ResponseWriter, req *http.Request) {
     fmt.Println("Requets Method in foo is:", req.Method, "\n")
     tpl.ExecuteTemplate(w, "index.gohtml", nil)
}


