package main

import (
        "net/http"
        "log"
        "flag"
)

var (
      listen = flag.String("listen", ":8080", "listen address")
      dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
     flag.Parse()
     log.Printf("Listening on %q...", *listen)
     http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))) 
     println("Hello World")
}
