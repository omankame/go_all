package main

import (
        "io"
        "net/http"
)

func main() {

//     mux := http.NewServeMux()
     http.HandleFunc("/", df)
     http.HandleFunc("/dog/", d)
     http.HandleFunc("/me/", o)
     
     http.ListenAndServe(":8080", nil)
}

func df(w http.ResponseWriter, req *http.Request) {
     io.WriteString(w, "WELCOME TO DEFAULT SERVER MUX")
}

func d(w http.ResponseWriter, req *http.Request) {
     io.WriteString(w, "DOG DOG DOG")
}

func o(w http.ResponseWriter, req *http.Request) {
     io.WriteString(w, "<h1> ONKAR </h1>")
}
