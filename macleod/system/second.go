package main

import (
        "os"
        "fmt"
//        "path/filepath"
)

func main() {
 
     args := os.Args[1:]
     if len(args) != 1 {
     fmt.Println("Please enter the filename")
     return
     }

     file := args[0]
     finfo, err := os.Stat(file)
     if err != nil {
     fmt.Println(err)
     }

    mode := finfo.Mode()
    fmt.Println(mode)

}   

    
