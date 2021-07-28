package main

import (
        "io"
        "net/http"
        "text/template"
)

var tpl *template.Template

func main() {

       tpl, _ = template.ParseFiles("something.gohtml") 


//     mux := http.NewServeMux()
     http.Handle("/", http.HandlerFunc(df))
     http.Handle("/dog/", http.HandlerFunc(d))
     http.Handle("/me/", http.HandlerFunc(o))
     
     http.ListenAndServe(":8080", nil)
}

func df(w http.ResponseWriter, req *http.Request) {
     io.WriteString(w, "WELCOME TO SECOND DEFAULT MUX")
}

func d(w http.ResponseWriter, req *http.Request) {
     io.WriteString(w, "SECOND DOG DOG DOG")
}

func o(w http.ResponseWriter, req *http.Request) {
     tpl.Execute(w, "ONKAR")
}
