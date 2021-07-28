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
     go Read(conn)
     go Write(conn)
}

func Read(conn net.Conn) {
     i := 0
     scnr := bufio.NewScanner(conn)
     if i == 0 {
        for scnr.Scan() {
        ln := scnr.Text()
        fmt.Println(ln)
        if ln == "" {
        break
        }
      }
    
    } else {
      break
    }
    i++ 
    fmt.Println("Code got here.")
}

func Write(conn net.Conn) {
     io.WriteString(conn, "I see you connected.\n")
}
