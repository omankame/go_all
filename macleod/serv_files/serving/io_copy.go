package main

import (
        "io"
        "os"
        "net/http"
)

func main() {

     http.HandleFunc("/", dog)
     http.HandleFunc("/toby.jpg", dogpic)
     http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, res *http.Request) {
     w.Header().Set("Content-Type", "text/html; charset=UTF-8")
     io.WriteString(w, `
        <img src="/toby.jpg">
        `)
}

func dogpic(w http.ResponseWriter, res *http.Request) {
     f, err := os.Open("toby.jpg")
     if err != nil {
        http.Error(w , "file not found", 404)
        return
     }
   
     defer f.Close()

     io.Copy(w, f)
//     if err != nil {
     

}
