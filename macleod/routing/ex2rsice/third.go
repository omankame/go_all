package main

import (
        "fmt"
        "os"
        "log"
        "net"
        "bufio"
        "io"
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
        continue
        }
     go handle(conn)

}

}

func handle(conn net.Conn) {
     defer conn.Close()
     scnr := bufio.NewScanner(conn)
     for scnr.Scan() {
        ln := scnr.Text()
        fmt.Println(ln)
     if ln == "" {
        break
     }
    }
   fmt.Println("Code got here.")
   io.WriteString(conn, "I see you connected.\n")
      
}
