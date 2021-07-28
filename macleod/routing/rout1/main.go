package main

import (
       "fmt"
       "net/http"
       "io"
)

type hotdog int

func (m hotdog)ServeHTTP(w http.ResponseWriter, req *http.Request) {

     switch req.URL.Path {
     case "/dog":
     io.WriteString(w , "DOG DOD DOG")
     case "/cat":
     fmt.Fprint(w, "CAT FMT CAT")
     default:
       fmt.Fprint(w, "THIS IS ROUTE DEMO")
   }      

}

func main() {

    var d hotdog
    http.ListenAndServe(":8080", d) 
}
