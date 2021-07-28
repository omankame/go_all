package main

import (
         "net/http"
         "cms3"
)

func main() {
     http.HandleFunc("/", cms3.ServeIndex)
     http.HandleFunc("/page/", cms3.ServePage)
     http.HandleFunc("/post/", cms3.ServePost)
     http.HandleFunc("/new", cms3.HandleNew)
     http.ListenAndServe(":3000", nil)


}
 
