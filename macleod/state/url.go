package main

import (
        "fmt"
        "net/http"
)

func main() {
     http.HandleFunc("/", foo)
     http.Handle("/favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
        v := req.FormValue("q")
        fmt.Fprint(w, "Do my search: ", v)   

}
