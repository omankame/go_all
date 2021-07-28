package main

import (
        "io"
        "os"
        "net/http"
//        "time"  
)

func main() {

     http.HandleFunc("/", dog)
     http.HandleFunc("/to", dogpic)
     http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
      w.Header().Set("Content-Type", "text/html; charset=UTF-8")
      io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogpic(w http.ResponseWriter, req *http.Request){

        f, err :=  os.Open("toby.jpg")
        if err != nil {
           http.Error(w , "file not found", 404)  
        }

        defer f.Close()

       fi, _ := f.Stat()

       http.ServeContent(w , req , f.Name() , fi.ModTime(), f)
}
