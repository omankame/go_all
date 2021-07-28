package main

/*  https://gobyexample.com/base64-encoding#:~:text=Example%3A%20Base64%20Encoding-,Go%20by%20Example%3A%20Base64%20Encoding,support%20for%20base64%20encoding%2Fdecoding.&text=This%20syntax%20imports%20the%20encoding,save%20us%20some%20space%20below.&text=Here's%20the%20string%20we'll%20encode%2Fdecode.  */

import (
        "fmt"
        "net/http"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
        "encoding/json"
//        "golang.org/x/crypto/bcrypt"
        
       b64 "encoding/base64"        
//        "html/template"
  

)


const (
        JSON = "application/json"
        TEXT = "application/x-www-form-urlencoded"
)


type Credentials struct {
        Password string `json:"password", db:"password"`
        Username string `json:"username", db:"username"`
}


func signup(w http.ResponseWriter, req *http.Request) {
     db = initDB()
     defer db.Close()
     creds := &Credentials{}

     if http.MethodPost == "POST" {    
        contentType := req.Header.Get("Content-type")
        switch contentType   {
        case JSON:
             fmt.Println("Onkar in")
             err := json.NewDecoder(req.Body).Decode(creds)
             if err != nil {
                fmt.Println("err")
                w.WriteHeader(http.StatusBadRequest)
                return
             }
             
 
            pwenc :=  b64.StdEncoding.EncodeToString([]byte(creds.Password))

 
            insUsr, err :=  db.Prepare("INSERT INTO Users(username, password  ) VALUES(?,?)")
            if err != nil {
            fmt.Println("Database issue")
            w.WriteHeader(http.StatusInternalServerError)
            return
            }

            _, err =  insUsr.Exec(creds.Username, pwenc)
            if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
           }

           json.NewEncoder(w).Encode(msg1)
           return

        case TEXT:
          req.ParseForm()
          uname := req.FormValue("username")
          pwd   := req.FormValue("password")

          pwenc :=  b64.StdEncoding.EncodeToString([]byte(pwd))

          insU, err :=  db.Prepare("INSERT INTO Users(username, password  ) VALUES(?,?)")
          if err != nil {
          fmt.Println("Database issue")
          w.WriteHeader(http.StatusInternalServerError)
          return
          }

           _, err =  insU.Exec(uname, pwenc)
           if err != nil {
              db.Close()
              w.WriteHeader(http.StatusInternalServerError)
              return
          }
           tpl.ExecuteTemplate(w, "signup_page.gohtml", msg1)
           return
 


        }
     }
      tpl.ExecuteTemplate(w, "signup_page.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
     db = initDB()
     defer db.Close()
//     creds := &Credentials{}
  
     if http.MethodPost == "POST" {
        contentType := req.Header.Get("Content-type")

         switch contentType   {
             case JSON:
             creds := &Credentials{}
             err := json.NewDecoder(req.Body).Decode(creds)
             if err != nil {
              w.WriteHeader(http.StatusBadRequest)
              return
             }
            
           result := db.QueryRow("select password from Users where username = ? ", creds.Username) //get password
           storedCreds := &Credentials{}

           // Store the obtained password in `storedCreds`
           err = result.Scan(&storedCreds.Password)
           if err != nil {
                // If an entry with the username does not exist, send an "Unauthorized"(401) status
                if err == sql.ErrNoRows {
                        w.WriteHeader(http.StatusUnauthorized)
                        return
                }
                // If the error is of any other type, send a 500 status
                fmt.Println(err, "lagta hey locchya")
                w.WriteHeader(http.StatusInternalServerError)
                return
           }

          sDec, _ := b64.StdEncoding.DecodeString(storedCreds.Password)
          strPwd := string(sDec)
          fmt.Println(strPwd)
          
            if strPwd == creds.Password  {
            json.NewEncoder(w).Encode(msg2)
            return    
            } else {
              
              w.WriteHeader(http.StatusBadRequest)
              return
            }      

          case TEXT:
           fmt.Println("Sign page inside")
           req.ParseForm()
           uname := req.FormValue("username")
           pwd   := req.FormValue("password")
           creds := &Credentials{}

           result := db.QueryRow("select password from Users where username = ? ", uname) //get password from the database

           err := result.Scan(&creds.Password)
           if err != nil {
           w.WriteHeader(http.StatusBadRequest)
           json.NewEncoder(w).Encode("The provided user does not exists in the database")
           return
           }


           sDec, _ := b64.StdEncoding.DecodeString(creds.Password)
           strPwd := string(sDec)
           fmt.Println(strPwd)

            if strPwd == pwd  {
//            json.NewEncoder(w).Encode(msg2)
            tpl.ExecuteTemplate(w, "sigin_page.gohtml", msg2)
            return
            } else {
              json.NewEncoder(w).Encode("Seems to be password provided is wrong") 
//              w.WriteHeader(http.StatusBadRequest)
              return
            }


         }
     }

     tpl.ExecuteTemplate(w, "sigin_page.gohtml", nil) 
}
 
