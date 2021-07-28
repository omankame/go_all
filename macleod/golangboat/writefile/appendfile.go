package main

import (
        "os"
        "fmt"
)

func main() {
     fl, err := os.OpenFile("onkar.txt", os.O_APPEND| os.O_WRONLY, 0644) 
     if err != nil {
        fmt.Println(err)
        return
     }

     defer fl.Close()

     line := "File append Exmplae"
 
     _, err = fmt.Fprintln(fl, line)
     
     fmt.Println("Operation sucessfull")

}
