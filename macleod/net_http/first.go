package main

import(
        "fmt"
        "net/http"
)

type hotdog int
// type Handler Interface {
//        func ServeHTTP(w ResponseWriter, req *Request)
//}

func (m hotdog)ServeHTTP(w http.ResponseWriter, req *http.Request) {
      fmt.Fprintln(w, "Any code you want to write in this function")
}

func main() {

     var d hotdog
     http.ListenAndServe(":8080", d)  
 

}
 
