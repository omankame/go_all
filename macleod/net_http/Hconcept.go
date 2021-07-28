package main

import (
        "fmt"
        "net/http"
)

type hotdog int

// type Handler interface {
//     ServerHttp(ResponseWriter, *request)

//} 

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
     fmt.Println("Any Code you want in this func")     
 
}

func main() {


}
