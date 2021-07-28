package main

import (
         "log"
         "bufio"
         "net"
         "fmt"
         "os"
         "strings"
)

func main() {
     li, err := net.Listen("tcp", ":8080")
     if err != nil {
     log.Fatalln(err)
     os.Exit(1)
     }

     defer li.Close() 

     for {
     conn ,err := li.Accept()
     if err != nil {
     log.Println(err)
     continue
     }

     go handle(conn)
     }
    
}

func handle(conn net.Conn) {

     defer conn.Close()

     request(conn)
     respond(conn)

}

func request(conn net.Conn) {
     i := 0
 
     scn := bufio.NewScanner(conn)
     for scn.Scan() {
         txt := scn.Text()
//         fmt.Fprintln(conn, txt)
         if i == 0 {
         m := strings.Fields(txt)[0]
         fmt.Println("***METHOD", m)
         }
        if txt  == "" {
            fmt.Println("Now Breaking")
            break
        }
       i++
 
       
     }   

}

func respond(conn net.Conn) {

        body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

       fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
        fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
        fmt.Fprint(conn, "Content-Type: text/html\r\n")
        fmt.Fprint(conn, "\r\n")
        fmt.Fprint(conn, body)


//            fmt.Fprint(conn, "\r\n")
  
//            fmt.Fprint(conn, "TESTING")
//            fmt.Fprint(conn, body)
}
            
