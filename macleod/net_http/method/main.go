package main

import (
        "log"
        "net/http"
        "html/template"
        "os"
//        "fmt"
        "net/url"
)

type hotdog int
var tpl *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//           fmt.Fprint(w, "Hi THIS is TESTING")
           err := req.ParseForm()
           if err != nil {
           log.Println(err)
           }

           data := struct {
                   Method      string
                   Submissions url.Values
                  }{
                     req.Method,
                     req.Form,
                 }  
           tpl.ExecuteTemplate(w, "index.gohtml", data)            

}

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
