package main

import (
//        "fmt"
        "net/http"
        "github.com/satori/go.uuid"
        "strings"
        "html/template"
)
var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
     http.HandleFunc("/", index)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
     c := getCookie(w, req)
     strData := getData()
//     c.Value = c.Value + "|"
     c.Value = c.Value + strData  
     slData := strings.Split(c.Value, "|")
   
     tpl.ExecuteTemplate(w, "index.gohtml", slData)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
     c, err := req.Cookie("session")
     if err != nil {
        id, _ := uuid.NewV4()
        c = &http.Cookie {
         Name: "session",
         Value: id.String(),
        }
      http.SetCookie(w, c)
    }
   return c
}
 
func getData() string{

     p1 := "disneyland.jpg"
     p2 := "abc.jpg"
     p3 := "reliance.jpg"

    var s string
    s = s + "|"
    if !(strings.Contains(s, p1)) {
       s =  s + p1 
    }
    
    if !(strings.Contains(s, p2)) {
       s = s + "|" 
       s = s + p2 
    }
    
   if !(strings.Contains(s, p3)) {
       s = s + "|" 
       s = s + p3 
    }

//    c , _  := req.Cookie("session")
//    s = "|" + c.Value
    return s
    
}


