package main

import (
        "fmt"
        "io"
        "net/http"
)

type hotdog int

func (d hotdog)ServeHTTP(w http.ResponseWriter, r *http.Request) {
            fmt.Fprint(w, "DOG DOG DOG")      
}
type hotcat int
func (c hotcat)ServeHTTP(w http.ResponseWriter, r *http.Request) {
           io.WriteString(w, "CAT CAT CAT")      
}

func main() {

     var d hotdog
     var c hotcat

     mux := http.NewServeMux()
     mux.Handle("/dog/", d)
     mux.Handle("/cat", c)
      
     http.ListenAndServe(":8080", mux)
}



