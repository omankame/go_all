package main

import (
         "fmt"
         "log"
         "bufio"
         "net"
)

func main() {

     li, err := net.Listen("tcp", ":80")
     if err != nil {
        log.Fatalln(err)
     }

    defer li.Close()
   
    for {
         conn, err := li.Accept()
         if err != nil {
           log.Println(err)
           continue
         }
      go handle(conn)
    }

}

func handle(conn net.Conn) {

     scanner := bufio.NewScanner(conn)
     for scanner.Scan() {
           ln := scanner.Text()
           fmt.Println(ln)
     }

     defer conn.Close()

    fmt.Println("Code Got here")

}
   


        
    
