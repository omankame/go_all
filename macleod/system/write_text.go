package main

import (
        "os"
        "fmt"
)

func main() {
     args := os.Args[1:]
     if len(args) != 1 {
     fmt.Println("please enter the file name")
     return
     }
    
     fname := args[0]

     fl, err := os.Create(fname)
     if err != nil {
        fmt.Println(err)
        os.Exit(1)
      } 

     defer fl.Close()

     fmt.Fprint(fl, "Test:") 
     fmt.Fprintln(fl, " Ok, let us see how it works")     
} 

