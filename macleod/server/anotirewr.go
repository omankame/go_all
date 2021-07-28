package main

import (
         "fmt"
         "log"
         "bufio"     
         "net"
         "time"
         "os"

)

func main() {
     
     li, err := net.Listen("tcp", ":80")
     if err != nil {
        log.Fatalln(err)
        os.Exit(1)
     }

//     t     := (20 * time.Second) 

       
   
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
     t := time.Now().Add(20 * time.Second) 
     err := conn.SetDeadline(t) 
     if err != nil {
                log.Fatalln("CONN TIMEOUT")
        }

     scanner := bufio.NewScanner(conn)

     for scanner.Scan() {
         txt := scanner.Text()
         fmt.Println(txt)
         fmt.Fprintln(conn, "Got it what you have wriiten")
       }

      defer conn.Close()
   
      fmt.Println("Code reached here")

}
     
