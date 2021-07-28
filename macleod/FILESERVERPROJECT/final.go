package main

import (
        "net/http"
        "net"
        "io"
//        "os"
        "fmt"
        "time"
)

func main() {
     http.HandleFunc("/", web)
     http.HandleFunc("/index", donwload)
     http.Handle("favicon.ico", http.NotFoundHandler())
     http.ListenAndServe(":8080", nil)
}

func web (w http.ResponseWriter, r *http.Request) {
     fmt.Fprint(w, "Welcome to Download Demo, GO to INdex PAge")
}

func donwload(w http.ResponseWriter, r *http.Request) {
     url := "https://upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png"

     timeout := time.Duration(5) * time.Second

     transport := &http.Transport {
                  TLSHandshakeTimeout: timeout,
                  Dial: func(network, addr string) (net.Conn, error) {
                        return net.DialTimeout(network, addr, timeout)
                       },
                  DisableKeepAlives: true,
     }

    client := &http.Client {
              Transport: transport,
    }

    resp, err := client.Get(url)
    if err != nil {
       fmt.Println(err)
    }
    defer resp.Body.Close()

    w.Header().Set("Content-Disposition", "attachment; filename=Wiki.png")
    w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

    io.Copy(w, resp.Body)

}
      




