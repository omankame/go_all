package main

import (
        "html/template"
        "net/http"
//        "log"
        "os"
)

var tpl *template.Template

func init() {
      tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
 
      http.Handle("/", dog)
      http.ListenAndServe(":8080", nil)     
      tpl.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
}   
     

       
