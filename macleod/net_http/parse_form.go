package main

import (
//        "fmt"
        "net/http"
        "html/template"
        "log"
        "os"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//      fmt.Fprintln(w, "HI THIS IS HTTP SERVER")  
      err := req.ParseForm()
      if err != nil {
      log.Fatal(err)
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
