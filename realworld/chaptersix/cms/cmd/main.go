package main

import (
        "net/http"
        "cms"
)


     
func main() {
     http.HandleFunc("/", cms.ServeIndex)
     http.HandleFunc("/page/", cms.ServePage)
     http.HandleFunc("/post/", cms.ServePost)
     http.HandleFunc("/new", cms.HandleNew)
     http.HandleFunc("/newcomment", cms.HandleComment)
     http.HandleFunc("/allcomment", cms.HandleAllComment)
     http.ListenAndServe(":3000", nil)

// allcomment badlun post related nav thev ani all commnet je ahe te post related comment sathi wapar


}

