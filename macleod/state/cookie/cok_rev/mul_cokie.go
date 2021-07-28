package main

import (
        "net/http"
        "fmt"
)

func main() {
     http.HandleFunc("/", set)
     http.HandleFunc("/read", read)
     http.HandleFunc("/abundance", abundance)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

     cookie := &http.Cookie {
                Name:  "onkar-cookie",
                Value: "OM",
                Path:	"/",
    }
 
    http.SetCookie(w , cookie)
    fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
    fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
     c, err := req.Cookie("onkar-cookie")
     if err != nil {
        http.Error(w , "Server did not responded", http.StatusInternalServerError) 
    }
    fmt.Fprintln(w, "Cookie: ", c)
  
   c1, err := req.Cookie("general")
   if err != nil {
      fmt.Println(err)
   }
   fmt.Fprintln(w, "Cookie: ", c1)

   c2, err := req.Cookie("specific")
   if err != nil {
      fmt.Println(err)
   }
   fmt.Fprintln(w, "Cookie: ", c2)
   
}

func abundance(w http.ResponseWriter, req *http.Request) {
     http.SetCookie(w ,  &http.Cookie {
                Name:  "general",
                Value: "GENERAL",
                Path:   "/",
                } )    

     
     http.SetCookie(w ,  &http.Cookie {
                Name:  "specific",
                Value: "SPECIFIC COKKIE",
                Path:   "/",
                } )    

}
   


