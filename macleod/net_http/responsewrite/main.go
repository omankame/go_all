package main

import ( 
        "fmt"
        "net/http"
//        "html/template"
//        "os"
)

type hotdog int

func (m hotdog)ServeHTTP(w http.ResponseWriter, r *http.Request) {
      
     w.Header().Set("ONKAR-KEY", "THIS IS FROM GOPHER ONKAR")
     w.Header().Set("Content-Type", "text/html; charset=utf-8")
     fmt.Fprint(w, "<h1> THE LOGIC YOU WANT CAN BE WRITE HERE TO DISPLAY </h1>")     

}

func main() {
 
     var d hotdog
     http.ListenAndServe(":8080", d)

}
