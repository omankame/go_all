package main

import (
        "time"
        "net"       
        "net/http"
        "io"
//        "io/ioutil"
//        "os"
)

func main() {
     http.HandleFunc("/", index)
     http.HandleFunc("/download", down)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
     io.WriteString(w, "Welcome to Main Page.")
}

func down(w http.ResponseWriter, req *http.Request) {
     url := "https://upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png"   
  
     timeout := time.Duration(5) * time.Second
 
     transport := &http.Transport {
                  ResponseHeaderTimeout: timeout,
                  Dial: func(network, addr string) (net.Conn, error) {
                       return net.DialTimeout(network, addr, timeout)
                      },
                  DisableKeepAlives: true,
                }

     client := &http.Client {
               Transport: transport,
             }

     resp, _ := client.Get(url)
     defer resp.Body.Close()

     //copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
	w.Header().Set("Content-Disposition", "attachment; filename=Wiki.png")
	w.Header().Set("Content-Type", req.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", req.Header.Get("Content-Length"))


    io.Copy(w, resp.Body)  
  
}
     
 
// imp notes := Dial specifies the dial function for creating unencrypted TCP connections. we need http transport and http client structure
// in transport need timeout and dial. dial use net package having func dialtimeout ==> act as dial with timeout option

