package main

import (
        "fmt"
        "net/http"
        "html/template"
        "github.com/satori/go.uuid"         
     )
var tpl *template.Template
var dbUsers = map[string]user{}  //userID, user
var dbSessions = map[string]string{} //session ID user ID

func init() {
     tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type user struct {
      UserName string
      First string
      Last string       
}

func main() {
     http.HandleFunc("/", submit) 
     http.HandleFunc("/bar", bar)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func submit(w http.ResponseWriter, req *http.Request) {
         
        id, _ := uuid.NewV4()
        cookie, err := req.Cookie("mysession")
        if err != nil {
           cookie = &http.Cookie {
                  Name:  "mysession",
                  Value: id.String(),
                  Path: "/",
                  HttpOnly: true,
               }
        http.SetCookie(w , cookie )
        }        
       fmt.Println(cookie) 
        var u user
        //if user already exists but i think no use
//        if un, ok := dbSessions[cookie.Value]; ok {
//            u = dbUsers[un]
//        }
        fmt.Println(u) 
        if req.Method == http.MethodPost {
           un := req.FormValue("username")
           f  := req.FormValue("firstname")
           l  := req.FormValue("lastname")
           u = user {
              	UserName: un,
              	First: f,
              	Last: l,
              }
           dbSessions[cookie.Value] = un
           dbUsers[un] = u
        }
        tpl.ExecuteTemplate(w , "index.gohtml", u)
    
}

func bar(w http.ResponseWriter, req *http.Request) {
     	c, err := req.Cookie("mysession")
      	if err != nil {
        http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
        }
     
        un, ok  := dbSessions[c.Value]
        if !ok {
            http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
        }
        u := dbUsers[un]
      tpl.ExecuteTemplate(w , "bar.gohtml", u)

}
