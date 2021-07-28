package main

import (
         "io"
         "net/http"
)

func c(res http.ResponseWriter, req *http.Request) {
      io.WriteString(res, "HANDLERFUNC CAT CAT CAT")
}

func d(res http.ResponseWriter, req *http.Request) {
      io.WriteString(res, "HANDLERFUNC DOG DOG DOG")
}

func df(w http.ResponseWriter, r *http.Request) {
     io.WriteString(w, "WELCOME TO ONKAR WEBSERVER THIS IS HANDLERFUNC EXAMPLE")
}
 


func main() {
    
       http.Handle("/", http.HandlerFunc(df)) 
       http.Handle("/cat", http.HandlerFunc(c))
       http.Handle("/dog/", http.HandlerFunc(d))
       http.ListenAndServe(":8080", nil) 

}
