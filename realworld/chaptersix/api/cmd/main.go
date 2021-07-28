package main

import (
         "net/http"
         "api"
)

func main() {
     http.HandleFunc("/", api.Doc)
     http.HandleFunc("/", api.AllPages)
     http.ListenAndServe(":3000", nil)

}

