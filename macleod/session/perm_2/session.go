package main

import (
        "net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
     var u user
     c, err := req.Cookie("session")
     if err != nil {
        return u
     }
     
     if un, ok := dbSessions[c.Value]; ok {
        u  = dbUSers[un]
     }
     return u
}

func alreadyLogIn(req *http.Request) bool {
     c, err := req.Cookie("session")
     if err != nil {
        return false
     }

    un := dbSessions[c.Value]
    _, ok := dbUSers[un]
    return ok
}



 
     
