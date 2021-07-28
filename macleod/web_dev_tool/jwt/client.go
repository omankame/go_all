package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "time"
        jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func main() {
     http.HandleFunc("/", homepage)
     http.ListenAndServe(":8081", nil)
}

func homepage(w http.ResponseWriter, req *http.Request) {
     validToken, err := genrateJWT()
     if err != nil {
        fmt.Println("Failed to generate Token")
        
     }
   
    client := &http.Client{}
    req, _  = http.NewRequest("GET", "http://10.32.138.72:8080/", nil)
    req.Header.Set("Token", validToken)
    res, err := client.Do(req)	
    if err != nil {
       fmt.Fprintf(w, "Error: %s", err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
       fmt.Println(err)
    }

    fmt.Fprintf(w, string(body))
                      
}

func genrateJWT() (string, error) {
     token := jwt.New(jwt.SigningMethodHS256)
 
     claims := token.Claims.(jwt.MapClaims)

     claims["authorized"] = true
     claims["client"] = "ONKAR M"
     claims["exp"] = time.Now().Add(time.Minute * 30).Unix() 

     tokenString, err := token.SignedString(mySigningKey)

     if err != nil {
        fmt.Errorf("Something went wrong: %s", err.Error())
        return "", err
     }

     return tokenString, nil

}
