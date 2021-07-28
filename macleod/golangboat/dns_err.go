package main

import (
        "fmt"
        "net"
)

func main() {
        addr, err :=  net.LookupHost("l.com")
        if err != nil {
              if dnErr, ok := err.(*net.DNSError);ok {
                 if dnErr.Temporary() {
                    fmt.Println("Temporary Error")
                    return
                  }
                  if dnErr.Timeout() {
                     fmt.Println("Timeoute Error")
                     return
                  }
                  fmt.Println("Generic DNS err")
                  return
               }
           fmt.Println("Generic Error")
       }

      fmt.Println("Address Values:", addr)
}
