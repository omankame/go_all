package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
)

func main() {
     http.HandleFunc("/hello/", foo)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)    
}

func foo(w http.ResponseWriter, req *http.Request) {
     switch req.Method {

     case "GET":
           for k, v := range req.URL.Query() {
               fmt.Printf("%s %s \n", k, v)
           }

          fmt.Fprint(w, "HTTP GET Recieved Requested")
     
     case "POST":
          body, err := ioutil.ReadAll(req.Body)
          if err != nil {
          fmt.Println(err)
          }
          fmt.Printf("%s\n", body)
          w.Write([]byte("Recieved POST REquest\n")) 
      
     default:
          w.WriteHeader(http.StatusNotImplemented)
          w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
    }

}
