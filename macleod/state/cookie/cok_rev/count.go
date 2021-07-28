package main

import (
        "net/http"
        "fmt"
        "strconv"
//      "io"
)

func main() {
     http.HandleFunc("/", set)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
     cookie, err := req.Cookie("my-cookie")
     if err == http.ErrNoCookie {
          cookie = &http.Cookie {
                Name:  "my-cookie",
                Value: "0",
                Path:  "/",
          }
    }

    count, _ := strconv.Atoi(cookie.Value)
    count++

    cookie.Value = strconv.Itoa(count)
    http.SetCookie(w , cookie )
//    io.WriteString(w , cookie.Value) 
   fmt.Fprint(w, "Website hit value:", cookie.Value)
}
