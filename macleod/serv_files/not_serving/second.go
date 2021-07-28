package main

import (
       "io"
       "net/http"
)

func main() {

     http.HandleFunc("/", dog)
     http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, res *http.Request) {
     
      w.Header().Set("Content-Type", "text/html; charset=UTF-8")

      io.WriteString(w , `
        <!--image doesn't serve-->
        <img src="/toby.jpg">
        `)
}
