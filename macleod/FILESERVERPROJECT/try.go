package main

import (
       "net/http"
       "fmt"
       "strings"
)

func main() {
     
     http.HandleFunc("/download", foo)
     http.HandleFunc("/", helloHandler)
     http.HandleFunc("/download/", zoo)
     http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintf(w, "Hello, %s\n", r.URL.Path[1:])
}

func foo(w http.ResponseWriter, r *http.Request) {
      fs := http.FileServer(http.Dir("/home/pdf/"))
      http.StripPrefix("/download", fs)      
}

func zoo(w http.ResponseWriter, r *http.Request) {
     str := r.URL.Path[1:]
//     fmt.Fprintf(w, "Hello, %s\n", str)
     sr := strings.Split(str, "/")    
     
     w.Header().Set("Content-Disposition", "attachment; filename=sr[1]")
     w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
     http.ServeFile(w, r, /home/pdf/sr[1])
}
