package main

import (
        "fmt"
        "os"
        "net"
)

func alreadyFileExists(fn string) bool {
   
     var _, err = os.Stat(fn)
     fmt.Println(err)
     if os.IsNotExist(err) {
        return false
     }

    return true

} 

func checkIPAddress(ip string)  bool {
        if net.ParseIP(ip) == nil {
           fmt.Printf("IP Address: %s - Invalid \n", ip)
           return false
        } else {
           fmt.Printf("IP Address: %s - Valid \n", ip)
           return true
       }

}   
