package main

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
        "fmt"
)

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

func Signup(w http.ResponseWriter, r *http.Request){
        db = initDB()
        defer db.Close()
	// Parse and decode the request body into a new `Credentials` instance
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	// Next, insert the username, along with the hashed password into the database

//	if _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword))
         insUsr, err :=  db.Prepare("INSERT INTO Users(username, password  ) VALUES(?,?)")
           if err != nil {
           fmt.Println("Database issue")
           w.WriteHeader(http.StatusInternalServerError)
           return
           }

          _, err =  insUsr.Exec(creds.Username, string(hashedPassword))
           if err != nil {
        
              w.WriteHeader(http.StatusInternalServerError)
              return
           }


	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
          json.NewEncoder(w).Encode(msg1) 
}

func Signin(w http.ResponseWriter, r *http.Request){
          db = initDB()
          defer db.Close()
  
	// Parse and decode the request body into a new `Credentials` instance	
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status		
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	// Get the existing entry present in the database for the given username
	result := db.QueryRow("select password from Users where username = ? ", creds.Username)
	if err != nil {
                 fmt.Println("onkar lochya hey re")
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We create another instance of `Credentials` to store the credentials we get from the database
	storedCreds := &Credentials{}
	// Store the obtained password in `storedCreds`
	err = result.Scan(&storedCreds.Password)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
                        fmt.Println("lochya")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
                fmt.Println(err, "lagta hey locchya")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
	}

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
        json.NewEncoder(w).Encode(msg2)
}
