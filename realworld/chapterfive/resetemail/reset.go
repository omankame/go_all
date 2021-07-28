package resetemail

import (
        "fmt"
        "crypto/rand"
        "encoding/base64"
        "html/template"
        "net/http"
        "net/smtp"    
//        "crypto/tls"     
//        "net"
)

var (
       validTokens []string
       servername = "smtp.mail.yahoo.com:465"
)

const passwordForm = `
<h1>Enter your new password</h1>
<form action="/reset" method="POST">
	<input type="hidden" name="email" value="{{ . }}" required>

	<label for="password">Password</label>
	<input type="password" name="password" required>

	<input type="submit" value="Submit">
</form>
`


func SendPasswordResetEmail(email string) error {
     token := string(genRandBytes())
     validTokens = append(validTokens, token)
      
     resetLink := "http://10.32.138.72:3333/reset?user=" + email + "&token=" + token
     fmt.Println(resetLink)
     username  := "Onkar.Mankame@ril.com"
     password  := "Reliance@2020"

//     host, _, _ := net.SplitHostPort("servername")
     
//      auth := smtp.PlainAuth("smtp.gmail.com:587", username, password, "smtp.gmail.com")
      auth := smtp.PlainAuth("smtp.mail.ril.com:587", username, password, "smtp.mail.ril.com")


    
    
    return smtp.SendMail("smtp.mail.ril.com:465", auth, username, []string{email}, []byte("Click here to reset your passsword: "+resetLink))  
//    return smtp.SendMail("smtp.gmail.com:587", auth, username, []string{email}, []byte("Click here to reset your passsword: "+resetLink))       
 
}

func genRandBytes() []byte {
     b := make([]byte, 24)
     _, err := rand.Read(b)
     if err != nil {
        panic(err)
     }

     return []byte(base64.URLEncoding.EncodeToString(b))
}

func ResetPassword(w http.ResponseWriter, req *http.Request) {

    switch req.Method {

     case "GET":
          var isValid bool
          var index int
 
           email := req.URL.Query().Get("user")
           token := req.URL.Query().Get("token")
 
           for i, tkn := range validTokens {
               if tkn == token {
                   isValid = true
                   index = i
                }
           }

           if !isValid {
              http.Error(w, "Token not valid", http.StatusUnauthorized)
           }

           validTokens = append(validTokens[:index], "")  // delete the token

           t  := template.Must(template.New("passwordForm").Parse(passwordForm))
           t.Execute(w, email)
           return
    
     case "POST":
  
          password := req.FormValue("password")
          email    := req.FormValue("email")

          req.ParseForm()

          err := OverrideOldPassword(email, password)
          if err != nil {
             http.Error(w, err.Error(), http.StatusInternalServerError)
          }
          
          w.Write([]byte("Sucessfully reset password"))

    }
 
}
