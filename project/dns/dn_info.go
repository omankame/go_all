package main

import (
        "fmt"
        "net"
        "os"
)

func main () {
     args := os.Args[1:]

     if len(args) != 1 {
        fmt.Println("please enter the url")
        return
     }

    ipChan := make(chan bool)
    cnChan := make(chan bool)
    go findIp(args[0], ipChan) 
    go findCname(args[0], cnChan)

    <-ipChan
    <-cnChan

    fmt.Println("All Done ':)' ")
}
     

func findIp(str string, ipChan chan bool) {
     ips, err := net.LookupIP(str)
     if err != nil {
        fmt.Println(err)
     }

      for _, v := range ips {
         fmt.Println(v)
      }

    ipChan <- true
}

func findCname(str string, cnChan chan bool) {
     cname, err := net.LookupCNAME(str)
      if err != nil {
        fmt.Println(err)
     }

     fmt.Println("Cname Value:", cname)
     cnChan <- true
} 
 
