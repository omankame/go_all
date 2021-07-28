package main

import (
        "fmt"
        "log"
        "net"
        "os"
        "io/ioutil"
)

func main() {
 
     conn, err := net.Dial("tcp", "localhost:80")  
     if err != nil {
        log.Fatalln(err)
        os.Exit(1)
     }

 
    bs, err := ioutil.ReadAll(conn)
    if err != nil {
       log.Println(err)
    }

    fmt.Printf("%s\n", bs) 


}
