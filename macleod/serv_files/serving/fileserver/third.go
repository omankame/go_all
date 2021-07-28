package main

import (
        "io"
        "net/http"
)

func main() {

     http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
     http.HandleFunc("/", dog)
     http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
         w.Header().Set("Content-Type", "text/html; charset=UTF-8")
         io.WriteString(w , `<img src="/assets/toby.jpg">`)
}
