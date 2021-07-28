package main

import (
         "net/http"
         "fmt"
         "io"
         "os"
           "github.com/gorilla/mux"
//         "path/filepath"
)

func main() {
     r := mux.NewRouter()

     r.Handle("/", http.FileServer(http.Dir("/home/macleod/FILESERVERPROJECT")))
     r.HandleFunc("/{name}", Download)
     http.ListenAndServe(":8080", r)
}

func Download(w http.ResponseWriter, req *http.Request) {
     url :=  "http://10.32.138.72:8080/download.go"

     param := mux.Vars(req) 
     fmt.Println(param)
// sample op of above is map[name:download.go]
//     var it string
//     for _, v := range param {
//        it = v
//     }

//     URL := filepath.Join(url, "/"+it)
//     fmt.Println(URL)
  
     resp, _ := http.Get(url)
     defer resp.Body.Close()  
     
    file, _ := os.Create("test")
    defer file.Close()
    
     w.Header().Set("Content-Disposition", "attachment; filename=test")
     w.Header().Set("Content-Type", req.Header.Get("Content-Type"))
     w.Header().Set("Content-Length", req.Header.Get("Content-Length"))

     io.Copy(file, resp.Body)  
}
     
