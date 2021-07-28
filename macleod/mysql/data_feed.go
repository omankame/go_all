package main

import (
        "net/http"
        "html/template"
        "fmt"
        "strconv"        
)

type Employee struct {
     Empid   	 int64
     First_name  string
     Last_name   string
}

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*")) 
}

func main() {
     http.HandleFunc("/", foo)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     
//e := req.FormValue("eid")
    e, _   := strconv.ParseInt(req.FormValue("eid"), 10, 64)
    f := req.FormValue("first")
    l := req.FormValue("last")
  
//    ei, _ := strconv.Atoi(e) 

    err :=  tpl.ExecuteTemplate(w, "index.gohtml", Employee{ e, f, l})
    if err != nil {
       fmt.Println(err)
      }

}



