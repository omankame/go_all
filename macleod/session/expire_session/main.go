package main

import (
//        "fmt"
        "net/http"
        "html/template"
        "github.com/satori/go.uuid"
        "time"
)

type user struct {
     UserName	string
     Password	[]byte
     First	string
     Last	string
     Role       string
}

type session struct {
     un			string
     lastActivity       time.Time
}

var tpl *template.Template
var dbSessions = map[string]session{}
var dbUsers = map[string]user{}
var dbsessionCleaned	time.Time

const sessionLength = 30

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
     dbsessionCleaned = time.Now()
}

func main() {
     http.HandleFunc("/", index)
     http.HandleFunc("/bar", bar)
     http.HandleFunc("/signup", signup)
     http.HandleFunc("/login", login)
     http.HandleFunc("/logout", logout)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
     u := getUser(w, req)
     
     tpl.ExecuteTemplate(w, "index.gohtml", u) 
}

func bar(w http.ResponseWriter, req *http.Request) {
    if !(alreadyLoggedIn(w, req)) {
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
   }
    c, _ := req.Cookie("session")
    s := dbSessions[c.Value]
    u := dbUsers[s.un]  
    tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
     if (alreadyLoggedIn(w, req)) {
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
     }
 
     if req.Method == http.MethodPost {
        un := req.FormValue("username")
        p := req.FormValue("password")
        f := req.FormValue("firstname")
        l := req.FormValue("lastname")
        r := req.FormValue("role")

        if _, ok := dbUsers[un]; ok {
        http.Error(w, "Username already exists", http.StatusForbidden)
        return
        }         

       pw, _ := GenerateFromPassword([]byte(p), bcrypt.MinCost) 
       u := user {un, pw, f, l, r}

       id, _ := uuid.NewV4()
       c := &http.Cookie {
            Name: "session",
            Value: id.String(),
            MaxAge: sessionLength,
        }
        http.SetCookie(w, c)
       dbSessions[c.Value] = session{un, time.Now()} 
       dbUsers[un] = u 
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
      }
      
    tpl.ExecuteTemplate(w, "signup.gohtml", nil)
} 

func login(w http.ResponseWriter, req *http.Request) {
     if (alreadyLoggedIn(w, req)) {
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
     }

     if req.Method == http.MethodPost {
        un := req.FormValue("username") 
        p := req.FormValue("password")
    
        u, ok := dbUsers[un]
        if !ok {
           http.Error(w, "Username and/or Password do not match", http.StatusForbidden)
           return
        }

        err := bcrypt.CompareHashAndPassword(u.Password,  []byte(p))
        if err != nil {
           http.Error(w, "Username and/or Password do not match", http.StatusForbidden)
           return
        }
       
        id, _ := uuid.NewV4()
        c := &http.Cookie {
             Name: "session",
             Value: id.String(),
             c.MaxAge: sessionLength,
        }
        http.SetCookie(w, c)
 
        dbSessions[c.Value] = session{un, time.Now()}
        dbUsers[un] = u
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
      }
   
         tpl.ExecuteTemplate(w, "login.gohtml", nil)
     } 

func logout(w http.ResponseWriter, req *http.Request) {
     if !(alreadyLoggedIn(w, req)) {
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
     }

     c, _ := req.Cookie("session")
     s := dbSessions[c.Value]
     delete(dbSessions, c.Value)
     c = &http.Cookie {
         Name: "session",
         Value: "",
         c.MaxAge: -1,
      }
      http.SetCookie(w, c)
      
      if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
         go cleanSessions()
      } 
      http.Redirect(w, req, "/login", http.StatusSeeOther)
} 
