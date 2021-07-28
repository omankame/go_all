package main

import (
        "net/http"
        "io"
        "crypto/sha256"
        "fmt"
        "crypto/hmac"
        "strings"
)

func main() {
     http.HandleFunc("/", foo)
     http.HandleFunc("/authenticate", validate)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     c, err := req.Cookie("onkar")
     if err != nil {
      c = &http.Cookie {
          Name: "onkar",
          Value: "",
       }
    }
    var fs string
    if req.Method == http.MethodPost {
       em := req.FormValue("email")
       hs := getHash(em)
       fs = em + "|" + hs
   }

   c.Value = fs
   http.SetCookie(w, c)
   io.WriteString(w, `<!DOCTYPE html>
        <html>
          <body>
            <form method="POST">
              <input type="email" name="email">
              <input type="submit">
            </form>
            <a href="/authenticate">Validate This `+c.Value+`</a>
          </body>
        </html>`)

}

func getHash(s string) string {
     hh := hmac.New(sha256.New, []byte("key"))
     io.WriteString(hh, s)
     return fmt.Sprintf("%x", hh.Sum(nil))
}

func validate(w http.ResponseWriter, req *http.Request) {
     c, err := req.Cookie("onkar")
     if err != nil {
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
     }
 
     if c.Value == "" {
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
      }
     sr := c.Value
     xs := strings.Split(sr, "|") 
     email := xs[0]
     codeRcvd := xs[1]
     codeCheck := getHash(email)
     if codeRcvd == codeCheck {
        io.WriteString(w, `<!DOCTYPE html>
        <html>
          <body>
                <h1>`+codeRcvd+` - RECEIVED </h1>
                <h1>`+codeCheck+` - RECALCULATED </h1>
          </body>
        </html>`)
      } else  {
        fmt.Println("Code does not match")
        fmt.Println(codeRcvd)
        fmt.Println(codeCheck)
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
     }     
}
     
