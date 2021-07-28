package main

import (
        "html/template"
        "net/http"
        "fmt"
)

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
     http.HandleFunc("/", foo)
     http.HandleFunc("/bar", barfoo)
     http.HandleFunc("/barred", boo)
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     fmt.Println("Request Method is: ", req.Method, "\n\n")  
}

func barfoo(w http.ResponseWriter, req *http.Request) {
     fmt.Println("Request Method is: ", req.Method, "\n\n") 
     http.Redirect(w,  req, "/", http.StatusSeeOther) 
}

func boo(w http.ResponseWriter, req *http.Request) {
     fmt.Println("Request Method is: ", req.Method, "\n")  
     tpl.ExecuteTemplate(w , "index.gohtml", nil)      
}
