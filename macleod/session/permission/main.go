package main

import (
//        "fmt"
         "net/http"
         "html/template"
         "github.com/satori/go.uuid"
         "golang.org/x/crypto/bcrypt"
)
type user struct {
    UserName 	string
    Password	[]byte
    First	string
    Last	string
    Role	string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}
func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
     http.HandleFunc("/", index)
     http.HandleFunc("/signup", signup)
     http.HandleFunc("/bar", bar)
     http.HandleFunc("/login", login)
     http.HandleFunc("/logout", logout)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
     u := getUser(req)
     tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
     u := getUser(req)
     if !alreadyLoggedIn(req) {
     http.Redirect(w, req, "/", http.StatusSeeOther)
     return 
     }
//     c, _ := req.Cookie("session")
//     un := dbSessions[c.Value]
//     u =  
     if u.Role != "007" { 
     http.Error(w , "Hey!, you are not allowed to enter into bar", http.StatusForbidden)
     return
     }
     
    tpl.ExecuteTemplate(w, "bar.gohtml", u)
}     

func signup(w http.ResponseWriter, req *http.Request) {
       
     if req.Method == http.MethodPost {
        un := req.FormValue("username")
        p  := req.FormValue("password")
        f  := req.FormValue("firstname")
        l  := req.FormValue("lastname")
        r  := req.FormValue("role")
   
        if _, ok := dbUsers[un]; ok {
           http.Error(w, "username is already exists", http.StatusForbidden)
        }

        bs, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost) 
        u := user { un, bs, f, l, r }
     
        id, _ := uuid.NewV4()
        c := &http.Cookie {
             Name: "session",
             Value: id.String(),
        }
        http.SetCookie(w, c)
        dbSessions[c.Value] = un
        dbUsers[un] = u
        
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
   } 

     tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
     if alreadyLoggedIn(req) {
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
     }    
//     var u user
     if req.Method == http.MethodPost {
        un := req.FormValue("username")
        p  := req.FormValue("password")
        
        u, ok := dbUsers[un]
        if !ok {
           http.Error(w, "username and password do not match", http.StatusForbidden)
           return
        }
    
       err :=  bcrypt.CompareHashAndPassword(u.Password, []byte(p)) 
       if err != nil {
          http.Error(w, "username and password do not match", http.StatusForbidden)
           return
        }
     
        id, _ := uuid.NewV4()
        c := &http.Cookie {
             Name: "session",
             Value: id.String(),
        }
        http.SetCookie(w, c)
        dbSessions[c.Value] = un
        dbUsers[un] = u
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return

    }

     tpl.ExecuteTemplate(w, "login.gohtml", nil)     
}

func logout(w http.ResponseWriter, req *http.Request) {
     if !alreadyLoggedIn(req) {
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
     }
     c, _ := req.Cookie("session") 
//     un :=  dbSessions[c.Value]
     delete(dbSessions, c.Value)    
     c = &http.Cookie {
          Name: "session",
          Value: "",
          MaxAge: -1,
      }
     http.SetCookie(w, c)
     http.Redirect(w, req, "/login", http.StatusSeeOther)

}


  
