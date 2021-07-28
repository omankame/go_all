package main

import (
         "net/http"
//         "html/template"
         "fmt"
)

func main() {
     http.HandleFunc("/", foo)
     http.HandleFunc("/bar", foobar)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     fmt.Println("The Requets method is: ", req.Method, "\n")
}

func foobar(w http.ResponseWriter, req *http.Request) {
     fmt.Println("The Requets method is: ", req.Method, "\n")
     http.Redirect(w , req, "/", 301)
}
