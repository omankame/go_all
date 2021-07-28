package main

import (
         "net/http"
         "fmt"
//         "database/sql"
       _ "github.com/go-sql-driver/mysql"
         "encoding/json"
         "golang.org/x/crypto/bcrypt"
)

const (
        JSON = "application/json"
        TEXT = "application/x-www-form-urlencoded"
)

func signup(w http.ResponseWriter, req *http.Request) {
     Db = initDB()
     defer Db.Close()
     creds := &Credentials{}
//     contentType := req.Header.Get("Content-type")

     if http.MethodPost == "POST" {
           fmt.Println("test")
           contentType := req.Header.Get("Content-type")
           fmt.Println(contentType) 
           switch contentType   {
           case JSON: 
           fmt.Println("in json") 
           err := json.NewDecoder(req.Body).Decode(creds)
           if err != nil {
           w.WriteHeader(http.StatusBadRequest)
           return
           }
         
//           pwd := []byte(creds.Password)
           hashPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 4)
           fmt.Println(hashPassword)
           if err != nil {
           w.WriteHeader(http.StatusInternalServerError)
           return
           }    
     
           insUsr, err :=  Db.Prepare("INSERT INTO Users(username, password  ) VALUES(?,?)")
     	   if err != nil {
           fmt.Println("Database issue")
           w.WriteHeader(http.StatusInternalServerError)
           return
           }      
    

           _, err =  insUsr.Exec(creds.Username, string(hashPassword))
           if err != nil {
              Db.Close()
              w.WriteHeader(http.StatusInternalServerError)
              return
           }  

 
           json.NewEncoder(w).Encode(msg1) 
//           tpl.ExecuteTemplate(w, "signup_page.gohtml", msg1)
           return
 
          case TEXT :
          fmt.Println("I am inside ")
          req.ParseForm()
          uname := req.FormValue("username")
          pwd   := req.FormValue("password")

          hashPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 4)
          fmt.Println(string(hashPassword))
          if err != nil {
          w.WriteHeader(http.StatusInternalServerError)
          return
          }


          insU, err :=  Db.Prepare("INSERT INTO Users(username, password  ) VALUES(?,?)")
          if err != nil {
          fmt.Println("Database issue")
          w.WriteHeader(http.StatusInternalServerError)
          return
          }

           _, err =  insU.Exec(uname, hashPassword)
           if err != nil {
              Db.Close()
              w.WriteHeader(http.StatusInternalServerError)
              return
          }
//           json.NewEncoder(w).Encode(msg1)
           tpl.ExecuteTemplate(w, "signup_page.gohtml", msg1)
           return

 
        }
         
         
      }
     
          tpl.ExecuteTemplate(w, "signup_page.gohtml", nil)
}

     
func signin(w http.ResponseWriter, req *http.Request) {
     Db = initDB()
     defer Db.Close()
     creds := &Credentials{}

     if http.MethodPost == "POST" {
        contentType := req.Header.Get("Content-type")

        switch contentType   {
           case JSON:
           fmt.Println("sign inside json")
           err := json.NewDecoder(req.Body).Decode(creds)
           if err != nil {
              w.WriteHeader(http.StatusBadRequest)
           return
           }

           result := Db.QueryRow("select password from Users where username = ? ", creds.Username) //get password
           storeCred := &Credentials{}

           err = result.Scan(&storeCred.Password)
           if err != nil {
           w.WriteHeader(http.StatusBadRequest)
           json.NewEncoder(w).Encode("The provided user does not exists in the database")
           return
           }
         
           hp, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 4)
           fmt.Println(string(hp) )
           err = bcrypt.CompareHashAndPassword([]byte(storeCred.Password), []byte(creds.Password))
           if err != nil {
           fmt.Println("password problem")
           w.WriteHeader(http.StatusUnauthorized)
           return
           }

           json.NewEncoder(w).Encode(msg2) 
           return 
      
           case TEXT:
           fmt.Println("Sign page inside")
           req.ParseForm()
           uname := req.FormValue("username")
           pwd   := req.FormValue("password")


           result := Db.QueryRow("select password from Users where username = ? ", uname) //get password from the database
//           storeCred := &Credentials{} //for storing the password

           err := result.Scan(&creds.Password)
           if err != nil {
           w.WriteHeader(http.StatusBadRequest)
           json.NewEncoder(w).Encode("The provided user does not exists in the database")
           return
           }
 
       
        
           err = bcrypt.CompareHashAndPassword([]byte(creds.Password), []byte(pwd))
           if err != nil {
           fmt.Println("password does not match")
           w.WriteHeader(http.StatusUnauthorized)
           json.NewEncoder(w).Encode("Provided password does not match")
           return
           }
         
           tpl.ExecuteTemplate(w, "sigin_page.gohtml", msg2)
           return
       }
    }
           tpl.ExecuteTemplate(w, "sigin_page.gohtml", nil)

}
              


