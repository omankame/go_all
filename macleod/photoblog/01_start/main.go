package main

import(
//       "fmt"
       "net/http"
       "html/template"
)
var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
     http.HandleFunc("/", index)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
} 

func index(w http.ResponseWriter, req *http.Request) {
     tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
