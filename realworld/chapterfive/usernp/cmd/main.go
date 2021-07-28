package main

import (
         "fmt"
//         "os"
         "usernp"
         
)

func main() {
     username, password := "omankame", "qwerty123"

     err := usernp.NewUser(username, password)
     if err != nil {
        fmt.Printf("Could not create user: %s\n", err.Error())
     }


     password = "onkar"
     err = usernp.AuthenticateUser(username, password)
     if err != nil {
        fmt.Printf("Could not authenticate user: %s\n", err.Error())
     } 

}	

     
