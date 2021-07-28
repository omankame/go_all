package main

import (
//        "log"
//        "fmt"
        "net/http"
        "github.com/gorilla/mux"
)

func main() {
   
     r := mux.NewRouter()
     r.HandleFunc("/", home)
     http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     switch req.Method {

     case "GET":
     w.WriteHeader(http.StatusOK)
     w.Write([]byte(`{"message": "get called"}`))

     case "POST":
     w.WriteHeader(http.StatusCreated)
     w.Write([]byte(`{"message": "Post called"}`))

     case "PUT":
     w.WriteHeader(http.StatusAccepted)
     w.Write([]byte(`{"message": "Put called"}`))

     case "DELETE":
     w.WriteHeader(http.StatusOK)
     w.Write([]byte(`{"message": "delete called"}`))     

     default:
     w.WriteHeader(http.StatusNotFound)
     w.Write([]byte(`{"message": "not found"}`))
   }

}
      
