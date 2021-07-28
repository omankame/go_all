package main

import (
         "fmt"
         "bufio"
         "net"
         "os"
         "log"
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
        conn, err := li.Accept()
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
     
}

func request(conn net.Conn) {
     i := 0
     scnr := bufio.NewScanner(conn) 
     for scnr.Scan() {
         txt := scnr.Text()
         if i == 0 {
              mux(conn, txt)
//            m := strings.Fields(txt)[0]
//            u := strings.Fields(txt)[1]
//           fmt.Println("***METHOD ", m)
//           fmt.Println("***URL", u)
          }
         if txt == "" {
          break
         }
      i++
}
}

func mux(conn net.Conn, ln string) {
            m := strings.Fields(ln)[0]
            u := strings.Fields(ln)[1]
            fmt.Println("***METHOD ", m)
            fmt.Println("***URL", u)

           if m == "GET" && u == "/" {
                   index(conn)
           } 
           if m == "GET" && u == "/about" {
                  about(conn)
           } 
           if m == "GET" && u == "/contact" {
                   contact(conn)
           }
           if m == "GET" && u == "/apply" {
                    apply(conn)
           }
           if m == "POST" && u == "/apply" {
                    applyProcess(conn)
           }
}

     
func index(conn net.Conn) {
        body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
        <strong>INDEX</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        </body></html>`
        fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
        fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
        fmt.Fprint(conn, "Content-Type: text/html\r\n")
        fmt.Fprint(conn, "\r\n")
        fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
        body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
        <strong>ABOUT</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        </body></html>`

        fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
        fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
        fmt.Fprint(conn, "Content-Type: text/html\r\n")
        fmt.Fprint(conn, "\r\n")
        fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {

        body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
        <strong>CONTACT</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        </body></html>`

        fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
        fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
        fmt.Fprint(conn, "Content-Type: text/html\r\n")
        fmt.Fprint(conn, "\r\n")
        fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {

        body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
        <strong>APPLY</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        <form method="POST" action="/apply">
        <input type="submit" value="apply">
        </form>
        </body></html>`

        fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
        fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
        fmt.Fprint(conn, "Content-Type: text/html\r\n")
        fmt.Fprint(conn, "\r\n")
        fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {

        body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
        <strong>APPLY PROCESS</strong><br>
        <a href="/">index</a><br>
        <a href="/about">about</a><br>
        <a href="/contact">contact</a><br>
        <a href="/apply">apply</a><br>
        </body></html>`

        fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
        fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
        fmt.Fprint(conn, "Content-Type: text/html\r\n")
        fmt.Fprint(conn, "\r\n")
        fmt.Fprint(conn, body)
}



