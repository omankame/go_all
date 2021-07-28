package main

import(
//       "fmt"
       "net/http"
       "html/template"
       "github.com/satori/go.uuid"
       "strings"
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
     c = appendValue(w, c) 
     xs := strings.Split(c.Value, "|") 
     tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
     c, err := req.Cookie("session")
     if err != nil {
        id, _ := uuid.NewV4()
        c = &http.Cookie {
          Name: "session",
          Value: id.String(),
          HttpOnly: true,
        }
     http.SetCookie(w, c)
     }
  
     return c
}
    
func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
 
        p1 := "disneyland.jpg"
        p2 := "atbeach.jpg"
        p3 := "hollywood.jpg"     
  
        s := c.Value

        if !(strings.Contains(s, p1)) {
           s += "|" + p1
        }
        if !(strings.Contains(s, p2)) {
           s += "|" + p2
        }
 
        if !(strings.Contains(s, p3)) {
           s += "|" + p3
        }

        c.Value = s
        http.SetCookie(w, c)
        return c
}
