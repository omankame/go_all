package main

import (
        "html/template"
        "log"
        "net/http"
        
        "chat"
)

var tpl = template.Must(template.ParseFiles("./index.html"))

func main() {
     go chat.DefaultHub.Start()

     http.HandleFunc("/", home)
     http.HandleFunc("/ws", chat.WSHandler)
     http.ListenAndServe(":3000", nil)

}

func home(w ResponseWriter, req *http.Request) {
     tpl.Execute(w, nil)

}

 
