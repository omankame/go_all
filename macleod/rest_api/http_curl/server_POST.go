package main

import (
        "net/http"
        "fmt"
        "io/ioutil"
        "strings"
)

func main() {

     url := "https://reqres.in/api/users"


     payload := strings.NewReader("name=test&job=teacher")

     req, _ := http.NewRequest("POST", url , payload) 

     req.Header.Add("content-type", "application/x-www-form-urlencoded")
     req.Header.Add("cache-control", "no-cache")

     res, _ := http.DefaultClient.Do(req)
  
     defer res.Body.Close()

     body, _ := ioutil.ReadAll(res.Body)
    
     fmt.Println(string(body))


}

// post method means posting some information to web server. Like form filling and submit button. so data send via http protocol to
// backend database server. Also take closer look at ex here information pass to "reqres.in" webserver and content write into it.
// so how clinet will get acknowledge it. here header content-type will come into picture. so different type contnt-type avilable/
// you can use anything ex. Content-Type: text/html; charset=UTF-8, Content-Type: multipart/form-data; boundary=something  :)
