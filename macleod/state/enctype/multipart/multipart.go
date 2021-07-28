//The Content-Length entity header indicates the size of the entity-body, in bytes, sent to the recipient.
//size of body pass to the http server via request 

package main

import (
        "html/template" 
        "fmt"
        "net/http"
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
      
     bs := make([]byte, req.ContentLength)
     req.Body.Read(bs)
     fmt.Println(req.ContentLength)
     fmt.Println(string(bs))

     body := string(bs)
 
     tpl.ExecuteTemplate(w, "index.gohtml", body)

}

