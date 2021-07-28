package main

import (
         "fmt"
         "strconv"
         "net/http"
)

func main() {

     http.HandleFunc("/", fibonnaci)
     http.ListenAndServe(":8080", nil)


}

func fibonnaci(w http.ResponseWriter, req *http.Request) {

     var i  int
     t1, t2  := 0, 1
     nextTerm := t1 + t2

     w.Write([]byte("Hello Testing of fibbonanci: \n"))   

     v := req.URL.Query().Get("fibnum")
     fmt.Printf("%[1]v  %[1]T \n", v)
 

     fnum, err := strconv.Atoi(v)
     if err != nil {
        fmt.Println("err")
        return
     }

           
     for i = 1; i <= fnum;   {
         i += 1
        fmt.Printf("%d, ", nextTerm)
        if i > fnum {
        fmt.Fprint(w, strconv.Itoa(nextTerm)) 
        } else {
          fmt.Fprint(w, strconv.Itoa(nextTerm) + ", " )
        }
        t1 = t2;
        t2 = nextTerm;
        nextTerm = t1 + t2;
    }

}

