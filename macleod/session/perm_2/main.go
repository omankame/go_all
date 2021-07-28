package main

import (
//        "fmt"        
        "net/http"
        "html/template"
        "golang.org/x/crypto/bcrypt"
        "github.com/satori/go.uuid"
)

type user struct  {
     UserName string
     Password []byte
     First    string
     Last     string
     Role     string
}

var tpl *template.Template
var dbUSers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
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
     tpl.ExecuteTemplate(w , "index.gohtml", u)      
}

func bar(w http.ResponseWriter, req *http.Request) {
     if !alreadyLogIn(req) {
         http.Redirect(w, req, "/", http.StatusSeeOther)
         return
     }
     u := getUser(w, req)
     if u.Role != "007" {
        http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
        return
        } 
     tpl.ExecuteTemplate(w , "bar.gohtml", u)
     
}

func signup(w http.ResponseWriter, req *http.Request) {
    if alreadyLogIn(req) {
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
    }

    if req.Method == http.MethodPost {
       un := req.FormValue("username")
       pw := req.FormValue("password")
       f  := req.FormValue("firstname")
       l  := req.FormValue("lastname")
       r  := req.FormValue("role")

       if _, ok := dbUSers[un]; ok {
          http.Error(w , "username already exists", http.StatusForbidden)
          return
       }
       p, _  := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost) 
       u := user { un, p, f, l, r}
       
       id, _ := uuid.NewV4()       
       c := &http.Cookie {
            Name: "session",
            Value: id.String(),
       }
       http.SetCookie(w, c)
       dbSessions[c.Value] = un
       dbUSers[un] = u
  
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
       }
     tpl.ExecuteTemplate(w , "signup.gohtml", nil)     

}
    
func login(w http.ResponseWriter, req *http.Request) {
     if alreadyLogIn(req) {
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
    }

    if req.Method == http.MethodPost {
       un := req.FormValue("username")
       pw := req.FormValue("password")

       u, ok := dbUSers[un]
       if !ok {
        http.Error(w, "username and/or password does not match", http.StatusForbidden)
        return
       }

       err := bcrypt.CompareHashAndPassword(u.Password, []byte(pw)) 
       if err != nil {
       http.Error(w, "username and/or password does not match", http.StatusForbidden)
        return
       }
   
       id, _ := uuid.NewV4()
       c := &http.Cookie {
            Name: "session",
            Value: id.String(),
        }
       http.SetCookie(w, c)
       dbSessions[c.Value] = un
       dbUSers[un] = u

       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
       }
     tpl.ExecuteTemplate(w , "login.gohtml", nil)


   
}      
     
func logout(w http.ResponseWriter, req *http.Request) {
     if !(alreadyLogIn(req)) {
       http.Redirect(w, req, "/", http.StatusSeeOther)
       return
     }
 
     c, _ := req.Cookie("session")
     delete(dbSessions, c.Value) 

     c = &http.Cookie {
         Name: "session",
         Value: "",
         MaxAge: -1,
      }
     http.SetCookie(w, c)
     http.Redirect(w, req, "/", http.StatusSeeOther)

}
    

    
    



     
     

