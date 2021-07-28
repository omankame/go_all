package main

import (
        "io"
        "net/http"
)


func d(w http.ResponseWriter, r *http.Request) {
      io.WriteString(w, "DOG DOG DOG ")
}

func c(w http.ResponseWriter, r *http.Request) {
      io.WriteString(w, "CAT CAT CAT")
}

func df(w http.ResponseWriter, r *http.Request) {
      io.WriteString(w, "HELLO WELCOME TO HANDLE FUNC")
}

func main() {


    http.HandleFunc("/dog/", d)
    http.HandleFunc("/cat", c)
    http.HandleFunc("/", df)

    http.ListenAndServe(":8080", nil)
}
