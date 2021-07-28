package main

import (
         "net/http"
         "erhandle"
)

func main() {
     
     logr := erhandle.CreateLog("../websrv")    

     http.Handle("/", erhandle.Time(logr, hello))
     http.Handle("/panic", erhandle.Recover(panicker))
     http.ListenAndServe(":3333", nil)

}

func hello(w http.ResponseWriter, req *http.Request) {
     w.Write([]byte("HELLO FROM ONKAR"))
}


func panicker(w http.ResponseWriter, req *http.Request) {
     panic(erhandle.ErrInvalidID)
}



