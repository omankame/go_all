package main

import (
        "fmt"
        "net/http"
        "github.com/satori/go.uuid"
)

func main() {
     http.HandleFunc("/", index)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
     cookie, err := req.Cookie("my-cookie")
     if err != nil {
       id, _ := uuid.NewV4()
//     id = id.String()
//     fmt.Println(id.String())

      cookie =  &http.Cookie {
            Name: "my-cookie",
            Value: id.String(),
            Path: "/",
            HttpOnly: true,
         }
      http.SetCookie(w , cookie)
    }

    fmt.Println(cookie)

}
 
