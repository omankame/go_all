package main

import (
         "fmt"
         "resetemail"
         "net/http"
)

func main() {

     username, password := "omkarmankame@gmail.com", "qwerty123"

     err := resetemail.CreateUser(username, password)
     if err != nil {
        fmt.Printf("Could not create user %s \n", err.Error())
     }

     err = resetemail.AuthenticateUser(username, password)
     if err != nil {
        fmt.Printf("Authentication failed for %s \n", err.Error())
     }

     fmt.Printf("Succesfully created and authenticated user \033[32m%s\033[0m\n", username)

     err = resetemail.SendPasswordResetEmail(username)
     if err != nil {
        fmt.Println(err)
     }

     http.HandleFunc("/reset", resetemail.ResetPassword)
     http.ListenAndServe(":3333", nil)
}
