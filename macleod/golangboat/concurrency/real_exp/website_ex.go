package main

import (
        "fmt"
        "net/http"

)

func main() {
     website := []string{
                   "https://stackoverflow.com/",
                   "https://github.com/",
                   "https://www.linkedin.com/",
		   "http://medium.com/",
		   "https://golang.org/",
                 }
 
     for _, web := range website {
        getWebsite(web)
     } 

}

func getWebsite(web string) {
     
        resp, err := http.Get(web)
           if err != nil {
            fmt.Println(err)
            return
          }

        str := resp.StatusCode
         if str == 200 {
           fmt.Println(web, "Website is live")
         }
   }

