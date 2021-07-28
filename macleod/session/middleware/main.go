package main

import (
//        "fmt"
        "html/template"
        "net/http"
        "github.com/satori/go.uuid"
)
 

type user struct {
     UserName    string
     Password	 string
     First	 string
     Last	 string
     Role	 string
}

var dbSessions =  map[string]string{}
var dbUsers =	 map[string]user{}
var tpl *template.Template

func init() {
      tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
     http.HandleFunc("/", index)
     http.HandleFunc("/signup", signup)
     http.HandleFunc("/bar", bar)
     http.HandleFunc("/login", login)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
     u := getUser(req)
    tpl.ExecuteTemplate(w, "index.gohtml", u)
}
func login(w http.ResponseWriter, req *http.Request) {
     if req.Method == http.MethodPost {
        un := req.FormValue("username")
        p  := req.FormValue("password")
     
        u, ok  := dbUsers[un]
        if !ok {
        http.Error(w, "Username and/or password do not match", http.StatusForbidden)
        }
        if u.Password != p {
        http.Error(w, "username and/or password do not match", http.StatusForbidden)
        return
        }

        id, _ :=  uuid.NewV4()
        c := &http.Cookie {
             Name: "session",
             Value: id.String(),
        }
        http.SetCookie(w, c)
        dbSessions[c.Value] = un
        http.Redirect(w, req, "/", http.StatusSeeOther)    
      }
      tpl.ExecuteTemplate(w, "login.gohtml", nil)    
}

func signup(w http.ResponseWriter, req *http.Request) {
     if req.Method == http.MethodPost {
        un := req.FormValue("username") 
        p  := req.FormValue("password")
        f  := req.FormValue("firstname")
        l  := req.FormValue("lastname")
        r  := req.FormValue("role")

        if _, ok := dbUsers[un]; ok {
           http.Error(w , "Username already exists, Plese sign with other name", http.StatusSeeOther)
        }

        u := user{un, p, f, l, r}
        id, _ := uuid.NewV4() 
 
        c := &http.Cookie {
              Name: "session",
              Value: id.String(),
        }   
        http.SetCookie(w, c)
        dbSessions[c.Value] = un
        dbUsers[un] = u    
        http.Redirect(w, req, "/", http.StatusSeeOther)
    }
       
     tpl.ExecuteTemplate(w, "signup.gohtml", nil )
}


func bar(w http.ResponseWriter, req *http.Request) {
/*    if  !alreadyLoggedIn(req) {
      http.Redirect(w , req , "/", http.StatusSeeOther)
    }  */
     u := getUser(req)
     tpl.ExecuteTemplate(w, "bar.gohtml", u)
}    



