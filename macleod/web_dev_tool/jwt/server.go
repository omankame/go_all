package main
import (
        "fmt"
        "net/http"
        jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func main() {
     http.Handle("/", isAuthorized(homepage))
     http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, req *http.Request) {
     fmt.Fprint(w, "Hello World!")
     fmt.Println("Endpoint Hit: HOME PAGE")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
     return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    
     if req.Header["Token"] != nil { 
        token, err := jwt.Parse(req.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
               if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                     return nil, fmt.Errorf("There was an erro")
               }
               return mySigningKey, nil
        })

        if err != nil {
           fmt.Fprintf(w, err.Error())
        }

        if token.Valid {
           endpoint(w, req)
        }

     } else {
       fmt.Fprintf(w, "Not Authorized")
     }

    })

}
