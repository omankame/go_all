package main

import (
        "fmt"
        "net/http"
        "html/template"
        "io"
        "os"
        "github.com/satori/go.uuid"
        "crypto/sha1"
        "strings"
        "path/filepath"
)

var tpl *template.Template

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
      http.HandleFunc("/", index)
//      http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
      http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
      http.Handle("favicon.ico", http.NotFoundHandler())
      http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
     c := getCookie(w, req)
     
     if req.Method == http.MethodPost {
        mf, mh, err := req.FormFile("nf")
        if err != nil {
           http.Error(w , err.Error(), http.StatusInternalServerError)
           return
        }
     defer mf.Close()
     h := sha1.New()
     ext := strings.Split(mh.Filename, ".")[1]
     fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
     wd, err := os.Getwd()
     if err != nil {
        fmt.Println(err)
     }

     path := filepath.Join(wd, "public", "pics", "fname" )
     pf, err := os.Create(path) 
     if err != nil {
        fmt.Println(err)
     }
     mf.Seek(0, 0)
     io.Copy(pf,mf)
     defer pf.Close()
  
     c = appendData(w, c, fname)     
     }
 
     xl := strings.Split(c.Value, "|") 
     tpl.ExecuteTemplate(w, "index.gohtml", xl[1:])
}

func appendData(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
     s := c.Value 
     if !(strings.Contains(s, fname)) {
        c.Value += "|" + fname
     }
     http.SetCookie(w, c)
     return c
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
      


