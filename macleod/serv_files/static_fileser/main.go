package main

import(
       "flag" 
       "net/http"
       "log"
)

func main() {
     
     port := flag.String("p", "8080", "port to be used")
     directory := flag.String("d", "/home/pdf/", "directory to be used") 
     flag.Parse()  
 
//     http.Handle("/", http.StripPrefix("*directory", http.FileServer(http.Dir(directory))))
//     http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
  
   http.Handle("/", http.FileServer(http.Dir(*directory)))
   log.Fatal(http.ListenAndServe(":"+*port, nil))
}


// github source https://gist.github.com/paulmach/7271283
