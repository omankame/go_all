package main

import (
        "fmt"
//        "log"
        "net/http"
        "html"
)

func main() {

    http.HandleFunc("/", foo)
    http.HandleFunc("/hi", zoo)
    http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func zoo(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Hi")
}
