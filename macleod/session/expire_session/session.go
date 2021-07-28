package main

import (
        "net/http"
        "github.com/satori/go.uuid"
         
)

func getUser(w http.ResponseWriter, req *http.Request) user {
     c , err := req.Cookie("session")
     if err != nil {
        id, _ := uuid.NewV4()
        c = &http.Cookie {
            Name: "session",
            Value: id.String(),
        }
        // refresh session 
        c.MaxAge = sessionLength
        http.SetCookie(w, c)
     }
     var u user 
     s, ok := dbSessions[c.Value]
     if ok {
        s.lastActivity = time.Now()
        dbSessions[c.Value] = s
        u, _ = dbUsers[s.un]
     }
     return u  
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
     c , err := req.Cookie("session") 
     if err != nil {
        return false
     }

     s, ok := dbSessions[c.Value]
     if ok {
                s.lastActivity = time.Now()
                dbSessions[c.Value] = s
        }

      _, ok = dbUsers[s.un]
      c.MaxAge = sessionLength
      http.SetCookie(w, c)
      return ok
}
     

func cleanSessions() {
     fmt.Println("BEFORE CLEAN") // for demonstration purposes
     showSessions() 
     for k, v := range dbSessions {
         if time.Now().Sub(v.lastActivity ) > time.Second * 30 {
         delete(dbSessions, k)
         }
     }
     dbSessionsCleaned = time.Now()
     fmt.Println("AFTER CLEAN") // for demonstration purposes
        showSessions()             // for demonstration purposes
}

func showSessions() {
     for k, v := range dbSessions {
         fmt.Println(k, v.un)
     }
     fmt.Println("")
}
   
