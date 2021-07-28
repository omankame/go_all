package main

import (
         "os"
         "net/http"
         "flag"
)

func main() {
     var dir string
     port := flag.String("port", "3000", "Port on which server run") 
     path := flag.String("path", "", "Path to serve")
     flag.Parse()
    if *path == "" {
        dir,_  = os.Getwd()
    } else {
         dir = *path

    }


    http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}  
    
