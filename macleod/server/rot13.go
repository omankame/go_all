package main

import ( 
        "net"
        "log"
        "bufio"
        "fmt"
        "strings"
)

func main() {

     ln, err := net.Listen("tcp", ":80")
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

     scanner := bufio.NewScanner(conn)
     for scanner.Scan() {
         txt := strings.ToLower(scanner.Text())
         bs := []byte(txt)
         r := rotate(bs)
//         fmt.Fprintln(conn, "Got it what have you written") 
         fmt.Fprintf(conn, "%s - %s\n\n", txt, r) 

     }
 
     defer conn.Close()

}

func rotate(txt []byte) []byte {

     r13 := make([]byte, len(txt))
     for i , v := range txt {
          if v <= 109 {
                r13[i] = v + 13
          } else {
                r13[i] = v - 13
                }
      }

     return r13

}
     




