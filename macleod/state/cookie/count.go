package main

import (
        "net/http"
//        "html/template"
        "fmt"
)

//const SameSite = iota 
const SameSiteDefaultMode SameSite = iota 
func main() {
     http.HandleFunc("/", foo)
     http.HandleFunc("/read", zoo)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     http.SetCookie(w, &http.Cookie {
                       Name: "my-cookie-onkar",
                       Value: "some-value-onkar",
                       Path:  "/",
                       SameSite: iota +1,
                      })

     fmt.Fprintln(w, "COOKIE-WRIITEN CHECK IN YOUR BROWSER")
     fmt.Fprint(w, "chrom ==> more tools ==> Developer tool ==> Cookies ==> click on hyperlink")  
     fmt.Println("The number of times website hit :", &http.Cookie.SameSite)
}

func zoo(w http.ResponseWriter, req *http.Request) {
     c, err := req.Cookie("my-cookie-onkar")
     if err != nil {
     http.Error(w , http.StatusText(400), http.StatusBadRequest)
     return
     }
     fmt.Fprintln(w, "YOUR COOKIE :", c) 
     
}

