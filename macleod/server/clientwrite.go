package main

import (
         "fmt"
         "net"
         "log"
         "os"
         "io"         
)

func main() {

     conn, err := net.Dial("tcp", "localhost:80")
     if err != nil {
        log.Fatalln(err)
        os.Exit(1)
      }
     defer conn.Close()
   
 
    io.WriteString(conn, "\nHello from client and dialled \n")  
    fmt.Fprintln(conn, "ok thats great") 


}
