package main

import (
//        "fmt"
        "io"
        "net"
        "log"      
        "os" 
)

func main() {
     conn, err := net.Dial("tcp", ":8000")
     if err != nil {
        log.Fatalln(err)
     }

     go io.Copy(conn, os.Stdin)
     io.Copy(os.Stdout, conn)

}
