package main

import (
       "net/http"
       "fmt"
       "log"
)

func main() {
     http.HandleFunc("/", foo)
     http.HandleFunc("/read", zoo)
     http.HandleFunc("/abundance", abb)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     http.SetCookie(w , &http.Cookie {
                    	Name:  "onkar-cookie",
        		Value: "onkar-data",
        		Path: "/",
                       })          
    fmt.Println(w, "Your cookie stored in browser")
    fmt.Print(w, "browser ==> more tools ==> dev tools ==> application ==> click on URL.") 
}
func zoo(w http.ResponseWriter, req *http.Request) {
     c1, err := req.Cookie("onkar-cookie")
     if err != nil {
        log.Println(err)
     } else {
       fmt.Fprintln(w, "Your Cookie#1 :", c1)
     }  

     c2, err := req.Cookie("general")
     if err != nil {
        log.Println(err)
     } else {
       fmt.Fprintln(w, "Your Cookie#2 :", c2)
     }  

     
     c3, err := req.Cookie("specific")
     if err != nil {
        log.Println(err)
     } else {
       fmt.Fprintln(w, "Your Cookie#1 :", c3)
     }  

}

func abb(w http.ResponseWriter, req *http.Request) {
     
      http.SetCookie(w , &http.Cookie {
                        Name:  "general",
                        Value: "general data ok",
                        })

      http.SetCookie(w , &http.Cookie {
                        Name:  "specific",
                        Value: "specific data ok",
                        })

      fmt.Println(w, "Your cookie stored in browser")
    fmt.Print(w, "browser ==> more tools ==> dev tools ==> application ==> click on URL.")
}
