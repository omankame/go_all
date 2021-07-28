package main

import (
//        "log"
//        "fmt"
        "net/http"
        "github.com/gorilla/mux"
)

func main() {
   
     r := mux.NewRouter()
     r.HandleFunc("/", get).Methods(http.MethodGet)
     r.HandleFunc("/", post).Methods(http.MethodPost)
     r.HandleFunc("/", put).Methods(http.MethodPut)
     r.HandleFunc("/", delete).Methods(http.MethodDelete)
     r.HandleFunc("/", notfound)     
     http.ListenAndServe(":8080", r)
}

func get(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     w.WriteHeader(http.StatusOK)
     w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     w.WriteHeader(http.StatusCreated)
     w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     w.WriteHeader(http.StatusAccepted)
     w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     w.WriteHeader(http.StatusOK)
     w.Write([]byte(`{"message": "delete called"}`))
}
    
func notfound(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     w.WriteHeader(http.StatusNotFound)
     w.Write([]byte(`{"message": "not found"}`))
}

