package main

import (
         "net/http"
          "cms2"
)

func main() {

     http.HandleFunc("/", cms2.ServeIndex)
     http.HandleFunc("/page/", cms2.ServePage)
     http.HandleFunc("/post/", cms2.ServePost)
     http.HandleFunc("/new", cms2.HandleNew)
     http.ListenAndServe(":3000", nil)
}


