package main

import(
//       "flag"
       "net/http"
//       "log"
       "fmt"
//       "net/url"
         "io"
)

func main() {

//     port := flag.String("p", "8080", "port to be used")
//     directory := flag.String("d", "/home/pdf/", "directory to be used")
//     flag.Parse()

     http.HandleFunc("/", index)
//   http.HandleFunc("/first.go/", download)
     http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*directory)))
}

func index(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Disposition", "attachment; filename=test")
     
        fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)

        err := Downloadfile(url, name)
        if err != nil {
        panic(err)
       }


}

//func download(w http.ResponseWriter, r *http.Request) 
//}

