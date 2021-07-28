//The Content-Length entity-header field indicates the size of the entity-body, in decimal number of OCTETs, sent to the recipient or, in the case of the HEAD method, the size of the entity-body that would have been sent had the request been a GET.

package main

import (
        "fmt"
        "html/template"
        "net/http"
//        "io"
)

var tpl *template.Template

type person struct {
     FirstName  string
     LastName   string
     Subscribed bool
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

    fmt.Println(req.ContentLength)
    
    bs := make([]byte, req.ContentLength)
    req.Body.Read(bs)
    body := string(bs)     

     tpl.ExecuteTemplate(w, "index.gohtml", body )

}
