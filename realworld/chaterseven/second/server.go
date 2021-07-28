package main

import (
        "net"
        "os"
        "fmt"
        "log"
//        "io"
        "bufio"
        "time"
)

var clients = []net.Conn{}

func main() {
     switch os.Args[1] {
     case "serve":
          getServe()

     default:
       fmt.Println("PLease enter serve only")

    }

}


func getServe() {
     ln, err := net.Listen("tcp", ":8000")
     if err != nil {
        log.Fatalln(err)
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
     clients = append(clients, conn)
     input := bufio.NewScanner(conn)
     for input.Scan() {
          for _, c := range clients {
           time.Sleep(1 * time.Second)
           c.Write( []byte(input.Text() + "\n"))
          }
     }
    conn.Close()
} 


