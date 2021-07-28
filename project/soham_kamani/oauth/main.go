package main

import (
         "net/http"
//         "log"
         "fmt"
         "os"
         "encoding/json"
)

const clientID = "1082f7e2bc2a98846358"
const clientSecret = "4930ed71b4623006a7327898fd3e45b7466d08a3"

func main() {

     fs := http.FileServer(http.Dir("public"))
     http.Handle("/", fs)

     // httpClient to make external HTTP request 
        httpClient := http.Client{}
 
    // create new redirect route
   
       http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, req *http.Request) {
           err  := req.ParseForm()
           if err != nil {
                 fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
                 w.WriteHeader(http.StatusBadRequest)
           }
  
           code := req.FormValue("code")
//             code := req.URL.Query().Get("code")
             fmt.Println(clientID, clientSecret, code)       
    //hit the github oauth point
 
   reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
       fmt.Println(reqURL)
    r, err := http.NewRequest(http.MethodPost, reqURL, nil)

    if err != nil {
                fmt.Println("mamala")
                 fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
                 w.WriteHeader(http.StatusBadRequest)
           }


    // as we want response in the form of json hence below line
       req.Header.Set("accept", "application/json")

      res, err := httpClient.Do(r)
       if err != nil {
                 fmt.Println("OM")
                 fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
                 w.WriteHeader(http.StatusBadRequest)
           }

      defer res.Body.Close()
 

    // parse the body in `OAuthAccessResponse` struct
    var t OAuthAccessResponse

     err = json.NewDecoder(res.Body).Decode(&t)
           
      if err != nil {
                 fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
                 w.WriteHeader(http.StatusBadRequest)
           }


     w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
     w.WriteHeader(http.StatusFound)

     })

     http.ListenAndServe(":8080", nil)

} 

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}
