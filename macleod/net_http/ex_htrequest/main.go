package main

import (
        "log"
        "net/http"
        "html/template"
        "os"
//        "fmt"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

//      fmt.Fprint(w, "THIS IS WEBPAGE)
      err := req.ParseForm()
      if err != nil {
         log.Println(err)
     }
     tpl.ExecuteTemplate(w, "index.gohtml", req.Form)   

}

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseFiles("index.gohtml"))
  
}

func main() {

     var d hotdog
     err := http.ListenAndServe(":8080", d)
     if err != nil {
     log.Fatalln(err)
     os.Exit(1)
     }
}


