package main

import (
        "log"
        "os"
        "net"
        "fmt"
        "bufio"
        "io"
        "time"
)

func main() {
 
     switch os.Args[1] {

      case "serve":
            getServe()
      case "conn":
           getConn()

      default:
          fmt.Println("Please enter serve and conn word only")
    }   

}

     

func getConn() {
     conn, err := net.Dial("tcp", ":8080")
     if err != nil {
       log.Println("Failed to connect.", err)
     }
     go io.Copy(conn, os.Stdin)
     io.Copy(os.Stdout, conn)
} 

func getServe() {
     ln, err := net.Listen("tcp", ":8080")
     if err != nil {
        log.Fatalln(err)
     }

     defer ln.Close()

     for  {
          conn , err := ln.Accept()
          if err != nil {
             log.Println(err)
             continue
          }
         go handle(conn)

     }

}

func handle(conn net.Conn) {

    input := bufio.NewScanner(conn)
    for input.Scan() {
        ln := input.Text()
        fmt.Println(ln)
        time.Sleep(1)
        fmt.Fprintf(conn, "I heard you say: %s\n", ln)
    }

    defer conn.Close()

}
          

