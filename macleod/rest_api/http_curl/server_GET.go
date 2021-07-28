//https://www.codershood.info/2017/06/25/http-curl-request-golang/

package main

import (
     "fmt"
     "net/http"
     "io/ioutil"
)

func main() {
     url := "https://reqres.in/api/users"

     req, _ :=  http.NewRequest("GET", url, nil) 

     req.Header.Add("cache-control", "no-cache")     
     
     res, _ := http.DefaultClient.Do(req)
      
     defer res.Body.Close()
 
     body, _ := ioutil.ReadAll(res.Body)
     
     fmt.Println(string(body))
}

// server get method meand read only. asking some information from http server ex. www.google.com
