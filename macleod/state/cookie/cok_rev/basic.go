package main

import (
        "net/http"
        "fmt"
)

func main() {
   
     http.HandleFunc("/", set)
     http.HandleFunc("/read", read)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
     cookie := &http.Cookie {
               Name: "my-cookie",
               Value: "some value",
               Path:  "/",
           }

     http.SetCookie(w , cookie )
     fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
     fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
    c, err :=  req.Cookie("my-cookie")
    if err != nil {
       http.Error(w , "Server did not responding", http.StatusInternalServerError)
     }
    fmt.Fprint(w, "Cokkie: ", c)       

}

    


 
