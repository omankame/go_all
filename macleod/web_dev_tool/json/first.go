package main

import (
        "fmt"
        "net/http"
        "encoding/json"
)

type person struct {
     Fname  string
     Lname  string
     Item   []string
}

func main() {

     http.HandleFunc("/", foo)
     http.HandleFunc("/marsh", marsh)
     http.HandleFunc("/encod", encod)
     http.Handle("/favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
     s := `<!DOCTYPE html>
                <html lang="en">
                <head>
                <meta charset="UTF-8">
                <title>FOO</title>
                </head>
                <body>
                You are at JIO
                </body>
                </html>`
     w.Write([]byte(s))
}

func marsh(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     p1 := person {
           Fname: "JAMES",
           Lname: "BOND",
           Item: []string{"Suit", "Gun", "Wry sense of humor"},
     }
    
     bt, err := json.Marshal(p1)
     if err != nil {
        fmt.Println(err)
     }

     w.Write(bt)

}

func encod(w http.ResponseWriter, req *http.Request) {
     w.Header().Set("Content-Type", "application/json")
     p1 := person {
           Fname: "JACK",
           Lname: "SPARROW",
           Item: []string{"BOAT", "GUN", "Superb Actor"},
     }

     err := json.NewEncoder(w).Encode(p1) 
     if err != nil {
        fmt.Println(err)
     } 

}
