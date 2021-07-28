package main

import (
        "fmt"
        "log"
        "bufio"
        "net" 
        "strings"        
)

func main() {
     li, err := net.Listen("tcp", ":80")
     if err != nil {
     log.Fatalln(err)
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

     fmt.Fprint(conn, "\nIN-MEMORY DATABASE", "\n \nUSE:", "\n \t SET key value", "\n \t GET key", "\n \t DEL key", "\n \nEXAMPLE:", "\n \t SET fav chocolate", "\n \t GET fav \n \n")  
     
     
     inmem := make(map[string]string)

     scan := bufio.NewScanner(conn)
     for scan.Scan() {
     txt := scan.Text()
     sstr := strings.Fields(txt)    

       switch sstr[0] {
         case "SET":
               if len(sstr) != 3 {
                  fmt.Fprintln(conn, "PLEASE ENTER ONLY 3 VALUES")
                  continue
                }
               inmem[ sstr[1] ] = sstr[2] 
//               fmt.Fprintln(conn, inmem)              
         case "GET":
               v, found := inmem[ sstr[1] ]
               if found != true {
               fmt.Fprintln(conn, "Value not found")
               } else {
                 fmt.Fprintln(conn, v)
               }
                   
         case "DEL":
              k := sstr[1] 
              delete(inmem, k)
              
              fmt.Fprintln(conn, inmem) 
         default:
             fmt.Fprintln(conn, "INVALID COMMAND")
             
    }
      

     }
    

//     defer conn.Close()
}


