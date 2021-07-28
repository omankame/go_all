package main

import ( 
//       "fmt"
       "net"
       "io"
       "log"
       "os"
)

func main() {

     ln, err := net.Listen("tcp", ":8080")
     if err != nil {
        log.Fatalln(err)
        os.Exit(1)
     }

     defer ln.Close()

     for {
        conn, err := ln.Accept() 
        if err != nil {
        log.Println(err)
        }
     go handleConn(conn)
    }

}

func handleConn(conn net.Conn) {

     defer conn.Close()

     io.WriteString(conn, "I see you connected.\n")
}
