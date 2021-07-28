package main

import (
        "net"
        "fmt"
        "log"
        "bufio"
)

func main() {
     
     li , err := net.Listen("tcp", ":80")
      if err != nil {
         log.Fatalln(err)
         return
      }

     defer li.Close()

     for {
         conn, err := li.Accept() 
         if err != nil {
            log.Println(err)
            continue
         }

     go handle(conn)
    }
}

func handle(conn net.Conn) {
     scanner := bufio.NewScanner(conn)
     for scanner.Scan() {
         txt := scanner.Text()
         fmt.Println(txt)
       fmt.Fprintf(conn, "Server received what you said : %s \n", txt)
     }
     defer conn.Close()

    fmt.Println("Code  got here")

}
   
 

             
