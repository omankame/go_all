package main

import (
         "fmt"
         "log"
         "io"
         "net"
         "os"
)

func main() {

     li, err := net.Listen("tcp", ":80") 
     if err != nil {
       log.Fatalln(err)
       os.Exit(1)
    }

    defer li.Close()

   for { 
        conn, err := li.Accept()
        if err != nil {
          log.Println(err)
          continue
       }

       io.WriteString(conn, "\nHello from TCP server\n")
       fmt.Fprintln(conn, "How is day?")
       fmt.Fprintf(conn, "%v", "Well! I hope\n")
       conn.Close()

   }

 

}

