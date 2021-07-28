package main

import (
        "os"
        "log"
        "net"
        "bufio"
        "fmt"
)

func main() {

        ln, err := net.Listen("tcp", ":80")
        if err != nil {
           os.Exit(1)
        }

        defer ln.Close()


        for {
        conn, err := ln.Accept()
        if err != nil {
           log.Println(err)
        }
        go handleConnection(conn)
     }
 

}

func handleConnection(conn net.Conn) {

     scnner := bufio.NewScanner(conn)
     
     for scnner.Scan() {
 
     txt := scnner.Text()
     fmt.Printf("%s\n", txt)

     }   

     defer conn.Close()
 

}
