package main

import (
        "net/smtp"
        "fmt"
)

func main() {

       username  := "onkarmankame@yahoo.com"
       password  := "Redhat@123@123"  
       
       email := "omkarmankame@gmail.com"
      auth := smtp.PlainAuth("smtp.mail.yahoo.com:465", username, password, "smtp.mail.yahoo.com")

     fmt.Println("AUth Executed")

     smtp.SendMail("smtp.mail.yahoo.com:465", auth, username, []string{email}, []byte("Click here to reset your passsword: "))

     fmt.Println("Sucess")

}
