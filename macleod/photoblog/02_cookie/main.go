package main

import(
//       "fmt"
       "net/http"
       "html/template"
       "github.com/satori/go.uuid"
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
     c := getCookie(w , req)    
     tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
     c, err := req.Cookie("session")
     if err != nil {
        id, _ := uuid.NewV4()
        c = &http.Cookie {
          Name: "session",
          Value: id.String(),
//          HttpOnly: true,
        }
     http.SetCookie(w, c)
     }
  
     return c
}
    
  
