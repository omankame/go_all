package main

import (
         "io"
         "net/http"
)

func main() {

     http.HandleFunc("/", dog)
     http.HandleFunc("/toby.jpg", dogpick)
     http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, res *http.Request) {
     w.Header().Set("Content-Type", "text/html; charset=UTF-8")
     io.WriteString(w, `<img src="toby.jpg"> `)

}

func dogpick(w http.ResponseWriter, res *http.Request) {
      http.ServeFile(w, res, "toby.jpg")
}
