package main

import (
        "fmt"
        "os"
        "log"
        "net"
        "bufio"
)

func main() {

      ln, err := net.Listen("tcp", ":80")
      if err != nil {
        log.Fatalln(err)
        os.Exit(1)
     }

    for {
        conn, err := ln.Accept()
        if err != nil {
           log.Println(err)                  
        }

        defer ln.Close()
     
        go handleConnection(conn)
    }

}

func handleConnection(conn net.Conn) {

     scanner := bufio.NewScanner(conn)
     for scanner.Scan() {
         txt := scanner.Text()
         fmt.Fprintf(conn, "%s \n", txt)
         fmt.Fprintln(conn, "This is what you have written and 'NET' server has received.")
         }
      defer conn.Close()
} 




