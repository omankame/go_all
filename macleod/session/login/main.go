package main

import (
//        "fmt"
        "net/http"
        "html/template"
        "github.com/satori/go.uuid"
        "golang.org/x/crypto/bcrypt"
)

type user struct {
     UserName	string
     Password	[]byte
     First	string
     Last	string
}


var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}         

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
     bs, _ := bcrypt.GenerateFromPassword( []byte("password"), bcrypt.MinCost)
     dbUsers["test@test.com"] = user{"test@test.com", bs, "James", "Bond"}
}

func main() {
     http.HandleFunc("/", index)
     http.HandleFunc("/signup", signup)
     http.HandleFunc("/login", login)
     http.HandleFunc("/bar", bar)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
     u := getUser(req)
     tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
     
     if req.Method == http.MethodPost {
     un := req.FormValue("username")
     p  := req.FormValue("password")
     f  := req.FormValue("firstname")
     l  := req.FormValue("lastname")

     pw, _  := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)     
     
     id, _ := uuid.NewV4()
     c := &http.Cookie{
          Name:  "session",
          Value: id.String(),
      }
      http.SetCookie(w, c)   
  
     u := user{ un, pw , f, l}
     dbSessions[c.Value] = un
     dbUsers[un] = u
     http.Redirect(w, req, "/", http.StatusSeeOther)
    }

    tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
     if req.Method == http.MethodPost {
     un := req.FormValue("username")
     p  := req.FormValue("password")

//     pw , _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
     
     u, ok  := dbUsers[un]
     if !ok {
        http.Error(w, "Username and/or password do not match", http.StatusForbidden)
     }
     
     err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
     if err != nil {
     http.Error(w, "Username and/or password do not match", http.StatusForbidden)
     }     
     id, err  := uuid.NewV4()
     c := &http.Cookie {
          Name:  "session",
          Value: id.String(),
       }
     http.SetCookie(w, c)
     dbSessions[c.Value] = un
 
     http.Redirect(w, req, "/", http.StatusSeeOther)
    }

    tpl.ExecuteTemplate(w, "login.gohtml", nil) 
    
}

func bar(w http.ResponseWriter, req *http.Request) {
     u := getUser(req)
     if !alreadyLoggedIn(req) {
     http.Redirect(w, req, "/", http.StatusSeeOther)
     }
     tpl.ExecuteTemplate(w, "bar.gohtml", u)
}



  
