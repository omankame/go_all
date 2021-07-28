package main

import ( 
      "fmt"
      "net/http"
//      "io"
//      "net/url"
)

func main() {

     http.HandleFunc("/", foo)
     http.HandleFunc("/query/", qoo)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     fmt.Fprint(w, `<h1> Welcome TO GET QUERY </h1>`)
}

func qoo(w http.ResponseWriter, req *http.Request) {
     for _, v := range req.URL.Query() {
     fmt.Fprint(w, "query was : ", v)
     }
}

 
          

