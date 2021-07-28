package main

import (
        "fmt"
//        "log"
        "net/http"
        "html"
        "sync"
        "strconv"
)

var (
     mutex sync.Mutex
     counter int
)
func main() {

    http.HandleFunc("/", foo)
    http.HandleFunc("/hi", zoo)
    http.HandleFunc("/increment", inc)
    http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func zoo(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Hi")
}

func inc(w http.ResponseWriter, r *http.Request){
     mutex.Lock()
     counter++
     fmt.Fprint(w, strconv.Itoa(counter))
     mutex.Unlock()
}

