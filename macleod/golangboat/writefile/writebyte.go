package main

import (
        "fmt"
        "os"
)

func main() {
     fl, err := os.Create("test.txt")
     if err != nil {
        fmt.Println(err)
        return
     }

     defer fl.Close()
     s := "onkar\n"
     b  := []byte(s)
     n, err := fl.Write(b)
     if err != nil {
        fmt.Println(err)
        return
     }
 
     fmt.Println(n, "th bytes written sucessfully")
}
 
