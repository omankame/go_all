package main

import (
//        "log"
//        "fmt"
        "net/http"
        "github.com/gorilla/mux"
)

func main() {
   
     r := mux.NewRouter()
     api := r.PathPrefix("/api/v1").Subrouter()
     api.HandleFunc("", get).Methods(http.MethodGet)
     api.HandleFunc("", post).Methods(http.MethodPost)
     api.HandleFunc("", put).Methods(http.MethodPut)
     api.HandleFunc("", delete).Methods(http.MethodDelete)
     api.HandleFunc("", notfound)     
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

