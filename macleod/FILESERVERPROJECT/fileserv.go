package main

import (
       "net/http"
       "fmt"
       "github.com/gorilla/mux"
//       "strings"
       "os"
       "io"
)

const fsDir = "/home/test"
func main() {
        fs := http.FileServer( http.Dir( fsDir ))
        r  := mux.NewRouter()
        r.HandleFunc("/{name}", nameFile) 
        r.Handle("/", fs)
       
  
   
//  http.StripPrefix("/download", http.FileServer(http.Dir("./home/macleod/system"))))

//        log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/home/macleod/system"))))
        http.ListenAndServe(":8080", r)

}

func nameFile(w http.ResponseWriter, r *http.Request) {

//     fmt.Fprint(w, "WELCOME TO SITE for download refer http://:8080/download")

       path := r.URL.Path
//       path = strings.Replace(path, "/", "", -1)
       fmt.Printf("%[1]T %[1]v", path)
//       fmt.Fprint(w, "This is a valid name  ", path)
       fpath := "/home/test" + path
       fmt.Println(fpath)

       f, _ := os.Open(fpath)
       defer f.Close()
       io.Copy(w, f)

}
