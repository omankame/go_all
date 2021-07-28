package main

import (
         "io"
         "fmt"
         "net/http"
)

type hotdog int
type hotcat int
type hotdef int
func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
     io.WriteString(w, " FINAL DOG DOG DOG" )  
}

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
     fmt.Fprint(w, "FINAL CAT CAT CAT" )  
}
func (df hotdef) ServeHTTP(w http.ResponseWriter, req *http.Request) {
     fmt.Fprint(w, "FINAL WELCOME TO SERVERMUX MUX WEBSERVER" )  
}

func main() {
       var d hotdog
       var c hotcat
       var df hotdef
  
       http.Handle("/cat",  c)
       http.Handle("/dog/", d)
       http.Handle("/", df)           
       http.ListenAndServe(":8080", nil)

}
