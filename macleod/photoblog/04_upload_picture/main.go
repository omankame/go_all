package main

import (
        "html/template"
        "net/http"
        "os"
        "github.com/satori/go.uuid"
        "strings"
        "crypto/sha1"
        "path/filepath"
        "io"
        "fmt"
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
//     var fname string 
     if req.Method == http.MethodPost {
       mf, mh, err :=  req.FormFile("nf")
       if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
      }
     fmt.Println("\nfile:", mf, "\nheader:", mh, "\nerr", err) 
     defer mf.Close()
     ext := strings.Split(mh.Filename, ".")[1]
     h := sha1.New()
//     io.Copy(h, mf)
     fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
     wd, err := os.Getwd()
     if err != nil {
     fmt.Println(err)
     }
 
     path := filepath.Join(wd, "public", "pics", fname)
     nf, err := os.Create(path)
     if err != nil {
        fmt.Println(err)
     }
     defer nf.Close()
     mf.Seek(0,0)
     io.Copy(nf, mf)
     c = appendValue(w, c, fname)
   }
     xs := strings.Split(c.Value, "|")     
     tpl.ExecuteTemplate(w , "index.gohtml", xs) 
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
     c, err  := req.Cookie("session")
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

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
     
     s := c.Value
     if !(strings.Contains(s, fname)) {
        s += "|" + fname 
     }
     c.Value = s
     http.SetCookie(w, c)
     return c
}
