package main

import (
        "net/http"
)

func getUser(req *http.Request) user {
     var u user
     c, err  := req.Cookie("session") 
     if err != nil {
        return u
     }

     un := dbSessions[c.Value]
     u = dbUsers[un]
     return u
}

func alreadyLoggedIn(req *http.Request) bool {
     c, err := req.Cookie("session")
     if err != nil {
        return false
     }

     un := dbSessions[c.Value]
     _, ok := dbUsers[un]
     return ok
}
     
