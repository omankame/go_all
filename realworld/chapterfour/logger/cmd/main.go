package main

import (
        "logger"
        "net/http"
        
)

func main() {
     logr := logger.CreateLog("../section4") 

     http.Handle("/", logger.TIME(logr, hello))
     http.ListenAndServe(":3333", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
     w.Write([]byte("HELLO WORLD!"))
}

